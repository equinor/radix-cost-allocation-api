package models

import (
	"encoding/json"
	"fmt"
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

// Env instance variables
type Env struct {
	Context              string
	APIEnvironment       string
	ClusterName          string
	DNSZone              string
	UseLocalRadixApi     bool
	UseProfiler          bool
	DbCredentials        *DBCredentials
	SubscriptionCost     float64
	SubscriptionCurrency string
	Whitelist            *costModels.Whitelist
	Cluster              string
}

// NewEnv Constructor
func NewEnv() *Env {
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	var (
		context          = os.Getenv("RADIX_CLUSTER_TYPE")
		apiEnv           = os.Getenv("RADIX_ENVIRONMENT")
		clusterName      = os.Getenv("RADIX_CLUSTERNAME")
		dnsZone          = os.Getenv("RADIX_DNS_ZONE")
		subCost          = os.Getenv("SUBSCRIPTION_COST_VALUE")
		subCurrency      = os.Getenv("SUBSCRIPTION_COST_CURRENCY")
		whiteList        = os.Getenv("WHITELIST")
		cluster          = os.Getenv("RADIX_CLUSTER_NAME")
		useLocalRadixApi = envVarIsTrueOrYes(os.Getenv("USE_LOCAL_RADIX_API"))
		useProfiler      = envVarIsTrueOrYes(os.Getenv("USE_PROFILER"))
	)
	if context == "" {
		log.Error("'Context' environment variable is not set")
	}
	if apiEnv == "" {
		log.Error("'API-Environment' environment variable is not set")
	}
	if clusterName == "" {
		log.Error("'Cluster' environment variables is not set")
	}
	if dnsZone == "" {
		log.Error("'DNS Zone' environment variables is not set")
	}
	subscriptionCost, err := strconv.ParseFloat(subCost, 64)
	if err != nil {
		subscriptionCost = 0.0
		log.Info("Subscription Cost is invalid or is not set.")
	}
	if len(subCurrency) == 0 {
		log.Info("Subscription Cost currency is not set.")
	}

	list := &costModels.Whitelist{}
	err = json.Unmarshal([]byte(whiteList), list)

	if err != nil {
		log.Info("Whitelist is not set")
	}

	return &Env{
		Context:              context,
		APIEnvironment:       apiEnv,
		ClusterName:          clusterName,
		DNSZone:              dnsZone,
		SubscriptionCost:     subscriptionCost,
		SubscriptionCurrency: subCurrency,
		Whitelist:            list,
		Cluster:              cluster,
		UseLocalRadixApi:     useLocalRadixApi,
		UseProfiler:          useProfiler,
		DbCredentials:        getDBCredentials(),
	}
}

func envVarIsTrueOrYes(envVar string) bool {
	return strings.EqualFold(envVar, "true") || strings.EqualFold(envVar, "yes")
}

func (env *Env) GetRadixAPIURL() string {
	if env.UseLocalRadixApi {
		return "localhost:3002"
	} else {
		return fmt.Sprintf("server-radix-api-%s.%s.%s", env.APIEnvironment, env.ClusterName, env.DNSZone)
	}
}

func (env *Env) GetRadixAPISchemes() []string {
	if env.UseLocalRadixApi {
		return []string{"http"}
	} else {
		return []string{"https"}
	}
}

func getDBCredentials() *DBCredentials {
	portEnv := os.Getenv("PORT")
	port := 1433

	if p, err := strconv.Atoi(portEnv); err == nil {
		port = p
	}

	return &DBCredentials{
		Server:   os.Getenv("SQL_SERVER"),
		Database: os.Getenv("SQL_DATABASE"),
		UserID:   os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     port,
	}
}
