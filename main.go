package main

import (
	"fmt"
	"github.com/equinor/radix-cost-allocation-api/api/cost"
	"github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"net/http"
	"os"
)

const clusternameEnvironmentVariable = "RADIX_CLUSTERNAME"

func main() {
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	fs := initializeFlagSet()

	var (
		port        = fs.StringP("port", "p", defaultPort(), "Port where API will be served")
		clusterName = os.Getenv(clusternameEnvironmentVariable)
	)

	log.Debugf("Port: %s\n", *port)
	log.Debugf("Cluster: %s\n", clusterName)

	parseFlagsFromArgs(fs)

	errs := make(chan error)
	go func() {
		log.Infof("API is serving on port %s", *port)
		err := http.ListenAndServe(fmt.Sprintf(":%s", *port), router.NewServer(clusterName, getControllers()...))
		errs <- err
	}()

	err := <-errs
	if err != nil {
		log.Fatalf("Radix cost allocation api server crashed: %v", err)
	}
}

func getControllers() []models.Controller {
	return []models.Controller{
		cost.NewApplicationController(),
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
