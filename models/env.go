package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Env instance variables
type Env struct {
	APIEnvironment   string
	ClusterName      string
	DNSZone          string
	UseLocalRadixApi bool
	UseProfiler      bool
	DbCredentials    *DBCredentials
	Whitelist        *Whitelist
	Cluster          string
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
		apiEnv           = os.Getenv("RADIX_ENVIRONMENT")
		clusterName      = os.Getenv("RADIX_CLUSTERNAME")
		dnsZone          = os.Getenv("RADIX_DNS_ZONE")
		whiteList        = os.Getenv("WHITELIST")
		cluster          = os.Getenv("RADIX_CLUSTER_NAME")
		useLocalRadixApi = envVarIsTrueOrYes(os.Getenv("USE_LOCAL_RADIX_API"))
		useProfiler      = envVarIsTrueOrYes(os.Getenv("USE_PROFILER"))
	)
	if apiEnv == "" {
		log.Error("'API-Environment' environment variable is not set")
	}
	if clusterName == "" {
		log.Error("'Cluster' environment variables is not set")
	}
	if dnsZone == "" {
		log.Error("'DNS Zone' environment variables is not set")
	}

	list := &Whitelist{}
	err := json.Unmarshal([]byte(whiteList), list)

	if err != nil {
		log.Info("Whitelist is not set. ", err)
	}

	return &Env{
		APIEnvironment:   apiEnv,
		ClusterName:      clusterName,
		DNSZone:          dnsZone,
		Whitelist:        list,
		Cluster:          cluster,
		UseLocalRadixApi: useLocalRadixApi,
		UseProfiler:      useProfiler,
		DbCredentials:    getDBCredentials(),
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
