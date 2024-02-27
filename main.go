package main

import (
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/repository"
	"github.com/equinor/radix-cost-allocation-api/service"

	"github.com/equinor/radix-cost-allocation-api/api/cost"
	"github.com/equinor/radix-cost-allocation-api/api/report"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/router"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	env, ctx, err := models.NewEnv()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error:\n%s\n\n", err.Error())
		os.Exit(3)
	}

	fs := initializeFlagSet()
	port := fs.StringP("port", "p", defaultPort(), "Port where API will be served")

	log.Debug().Msgf("Port: %s", *port)
	log.Debug().Msgf("Cluster: %s", env.ClusterName)

	parseFlagsFromArgs(fs)

	errs := make(chan error)

	authProvider := auth.NewAuthProvider(ctx, env.OidcIssuer, env.OidcAudience)
	radixAPIClient := radix_api.NewRadixAPIClient(env)
	costService := getCostService(env)

	go func() {
		log.Info().Msgf("API is serving on port %s", *port)
		addr := fmt.Sprintf(":%s", *port)
		handler := router.NewServer(env.ClusterName, env.OidcAllowedAdGroups, authProvider,
			cost.NewCostController(radixAPIClient, costService),
			report.NewReportController(costService))
		errs <- http.ListenAndServe(addr, handler)
	}()

	if env.UseProfiler {
		go func() {
			log.Info().Msgf("Profiler endpoint is serving on port 7070")
			errs <- http.ListenAndServe("localhost:7070", nil)
		}()
	}

	err = <-errs
	if err != nil {
		log.Fatal().Err(err).Msg("Radix cost allocation api server crashed")
	}
}

func getCostService(env *models.Env) service.CostService {
	return createContainerCostService(env)
}

func createContainerCostService(env *models.Env) service.CostService {
	gormdb, err := repository.OpenGormSqlServerDB(
		repository.GetSqlServerDsn(env.DbCredentials.Server, env.DbCredentials.Database, env.DbCredentials.UserID, env.DbCredentials.Password, env.DbCredentials.Port),
	)
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

func defaultPort() string {
	return "3003"
}
