package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api"

	_ "net/http/pprof"

	"github.com/equinor/radix-cost-allocation-api/api/cost"
	"github.com/equinor/radix-cost-allocation-api/api/report"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	env := models.NewEnv()
	fs := initializeFlagSet()
	port := fs.StringP("port", "p", defaultPort(), "Port where API will be served")

	log.Debugf("Port: %s\n", *port)
	log.Debugf("Cluster: %s\n", env.ClusterName)

	parseFlagsFromArgs(fs)

	errs := make(chan error)

	costRepository := models.NewSQLCostRepository(env.DbCredentials)
	defer costRepository.CloseDB()

	ctx := context.Background()
	authProvider := auth.NewAuthProvider(ctx)
	radixAPIClient := radix_api.NewRadixAPIClient(env)

	go func() {
		log.Infof("API is serving on port %s", *port)
		addr := fmt.Sprintf(":%s", *port)
		handler := router.NewServer(env.ClusterName, authProvider,
			cost.NewCostController(env, costRepository, radixAPIClient),
			report.NewReportController(env, costRepository))
		errs <- http.ListenAndServe(addr, handler)
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
