package main

import (
	"context"
	"fmt"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"net/http"
	"os"

	"github.com/equinor/radix-cost-allocation-api/api/cost"
	"github.com/equinor/radix-cost-allocation-api/api/report"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	_ "net/http/pprof"
)

const clusternameEnvironmentVariable = "RADIX_CLUSTERNAME"

func main() {
	env := models.NewEnv()
	fs := initializeFlagSet()

	var (
		port        = fs.StringP("port", "p", defaultPort(), "Port where API will be served")
		clusterName = os.Getenv(clusternameEnvironmentVariable)
	)

	log.Debugf("Port: %s\n", *port)
	log.Debugf("Cluster: %s\n", clusterName)

	parseFlagsFromArgs(fs)

	errs := make(chan error)

	costRepository := models.NewSQLCostRepository(env.DbCredentials)
	defer costRepository.CloseDB()

	ctx := context.Background()
	authProvider := auth.NewAuthProvider(ctx)
	radixAPIClient := radix_api.NewRadixAPIClient(env)

	go func() {
		log.Infof("API is serving on port %s", *port)
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", *port), router.NewServer(clusterName, authProvider, getControllers(env, costRepository, radixAPIClient)...))
	}()

	if env.UseProfiler {
		go func() {
			log.Infof("Profiler endpoint is serving on port 7070")
			errs <- http.ListenAndServe("localhost:7070", nil)
		}()
	}

	err := <-errs
	if err != nil {
		log.Fatalf("Radix cost allocation api server crashed: %v", err)
	}
}

func getControllers(env *models.Env, repo models.CostRepository, radixapi radix_api.RadixAPIClient) []models.Controller {
	return []models.Controller{
		cost.NewCostController(env, repo, radixapi),
		report.NewReportController(repo),
	}
}

func initializeFlagSet() *pflag.FlagSet {
	// Flag domain.
	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "DESCRIPTION\n")
		fmt.Fprintf(os.Stderr, "Radix cost allocation api server.\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		fs.PrintDefaults()
	}
	return fs
}

func parseFlagsFromArgs(fs *pflag.FlagSet) {
	err := fs.Parse(os.Args[1:])
	switch {
	case err == pflag.ErrHelp:
		os.Exit(0)
	case err != nil:
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		fs.Usage()
		os.Exit(2)
	}
}

func defaultPort() string {
	return "3003"
}
