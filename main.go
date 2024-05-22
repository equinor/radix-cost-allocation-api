package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/repository"
	"github.com/equinor/radix-cost-allocation-api/service"

	"github.com/equinor/radix-cost-allocation-api/api/cost"
	"github.com/equinor/radix-cost-allocation-api/api/report"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/router"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

const (
	defaultPort        = "3003"
	defaultMetricsPort = "9090"
	defaultProfilePort = "7070"
)

func main() {
	setupLogger()

	env, err := models.NewEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read environment configuration")
	}
	printInfo(env)

	fs := initializeFlagSet()
	port := fs.StringP("port", "p", defaultPort, "Port where API will be served")
	metricPort := fs.String("metrics-port", defaultMetricsPort, "The metrics API server port")
	parseFlagsFromArgs(fs)

	servers := []*http.Server{
		initializeServer(*port, env),
		initializeMetricsServer(*metricPort),
	}
	if env.UseProfiler {
		log.Info().Msgf("Initializing profile server on port %s", defaultProfilePort)
		servers = append(servers, &http.Server{Addr: fmt.Sprintf("localhost:%s", defaultProfilePort)})
	}

	startServers(servers...)
	shutdownServersGracefulOnSignal(servers...)
}

func printInfo(env *models.Env) {
	log.Debug().Msgf("Cluster: %s", env.ClusterName)

	info, _ := debug.ReadBuildInfo()
	log.Info().Str("version", info.Main.Version).Msg("Running")
}

func startServers(servers ...*http.Server) {
	for _, srv := range servers {
		go func() {
			log.Info().Msgf("Starting server on address %s", srv.Addr)
			if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				log.Fatal().Err(err).Msgf("Unable to start server on address %s", srv.Addr)
			}
		}()
	}
}

func shutdownServersGracefulOnSignal(servers ...*http.Server) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)
	s := <-stopCh
	log.Info().Msgf("Received %v signal", s)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	for _, srv := range servers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Info().Msgf("Shutting down server on address %s", srv.Addr)
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Warn().Err(err).Msgf("shutdown of server on address %s returned an error", srv.Addr)
			}
		}()
	}

	wg.Wait()
}

func initializeServer(port string, env *models.Env) *http.Server {
	log.Info().Msgf("Initializing API server on port %s", port)
	authProvider := auth.NewAuthProvider(context.Background(), env.OidcIssuer, env.OidcAudience)
	radixAPIClient := radix_api.NewRadixAPIClient(env)
	costService := getCostService(env)
	handler := router.NewHandler(env.ClusterName, env.OidcAllowedAdGroups, authProvider,
		cost.NewCostController(radixAPIClient, costService),
		report.NewReportController(costService))

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}
}

func initializeMetricsServer(port string) *http.Server {
	log.Info().Msgf("Initializing metrics server on port %s", port)
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.NewMetricsHandler(),
	}
}

func getCostService(env *models.Env) service.CostService {
	return createContainerCostService(env)
}

func createContainerCostService(env *models.Env) service.CostService {
	gormdb, err := repository.OpenGormSqlServerDB(env.DbCredentials.Server, env.DbCredentials.Database, env.DbCredentials.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db repository")
	}

	repo := repository.NewGormRepository(gormdb)
	return service.NewContainerCostService(repo, env.Whitelist.List)
}

func initializeFlagSet() *pflag.FlagSet {
	// Flag domain.
	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)
	fs.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "DESCRIPTION\n")
		_, _ = fmt.Fprintf(os.Stderr, "Radix cost allocation api server.\n")
		_, _ = fmt.Fprintf(os.Stderr, "\n")
		_, _ = fmt.Fprintf(os.Stderr, "FLAGS\n")
		fs.PrintDefaults()
	}
	return fs
}

func setupLogger() {
	level := os.Getenv("LOG_LEVEL")
	pretty, _ := strconv.ParseBool(os.Getenv("LOG_PRETTY"))

	var logWriter io.Writer = os.Stderr
	if pretty {
		logWriter = &zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}
	}

	logLevel, err := zerolog.ParseLevel(level)
	if err != nil || logLevel == zerolog.NoLevel {
		logLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	log.Logger = zerolog.New(logWriter).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &log.Logger
}

func parseFlagsFromArgs(fs *pflag.FlagSet) {
	err := fs.Parse(os.Args[1:])
	switch {
	case errors.Is(err, pflag.ErrHelp):
		os.Exit(0)
	case err != nil:
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		fs.Usage()
		os.Exit(2)
	}
}
