package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Env instance variables
type Env struct {
	APIEnvironment      string
	ClusterName         string
	DNSZone             string
	UseLocalRadixApi    bool
	UseProfiler         bool
	DbCredentials       *DBCredentials
	Whitelist           *Whitelist
	Cluster             string
	OidcIssuer          string
	OidcAudience        string
	OidcAllowedAdGroups []string
}

// NewEnv Constructor
func NewEnv() *Env {
	zerolog.DurationFieldInteger = true
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if envVarIsTrueOrYes(os.Getenv("PRETTY_PRINT")) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})
	}

	var (
		apiEnv              = os.Getenv("RADIX_ENVIRONMENT")
		clusterName         = os.Getenv("RADIX_CLUSTERNAME")
		dnsZone             = os.Getenv("RADIX_DNS_ZONE")
		whiteList           = os.Getenv("WHITELIST")
		cluster             = os.Getenv("RADIX_CLUSTER_NAME")
		useLocalRadixApi    = envVarIsTrueOrYes(os.Getenv("USE_LOCAL_RADIX_API"))
		useProfiler         = envVarIsTrueOrYes(os.Getenv("USE_PROFILER"))
		issuer              = os.Getenv("TOKEN_ISSUER")
		audience            = os.Getenv("TOKEN_AUDIENCE")
		allowedAdGroupsJson = os.Getenv("AD_REPORT_READERS")
	)
	if apiEnv == "" {
		log.Error().Msg("'API-Environment' environment variable is not set")
	}
	if clusterName == "" {
		log.Error().Msg("'Cluster' environment variables is not set")
	}
	if dnsZone == "" {
		log.Error().Msg("'DNS Zone' environment variables is not set")
	}
	if issuer == "" {
		log.Error().Msg("'TOKEN_AUDIENCE' environment variables is not set")
	}
	if audience == "" {
		log.Error().Msg("'TOKEN_AUDIENCE' environment variables is not set")
	}
	if allowedAdGroupsJson == "" {
		log.Error().Msg("'AD_REPORT_READERS' environment variables is not set")
	}
	var allowedGroups struct {
		List []string `json:"groups"`
	}
	err := json.Unmarshal([]byte(allowedAdGroupsJson), &allowedGroups)
	if err != nil {
		log.Error().Err(err).Msg("could not parse json for allowedADGroups")
	}

	list := &Whitelist{}
	err = json.Unmarshal([]byte(whiteList), list)
	if err != nil {
		log.Info().Err(err).Msg("could not parse json for WHITELIST")
	}

	return &Env{
		APIEnvironment:      apiEnv,
		ClusterName:         clusterName,
		DNSZone:             dnsZone,
		Whitelist:           list,
		Cluster:             cluster,
		UseLocalRadixApi:    useLocalRadixApi,
		UseProfiler:         useProfiler,
		DbCredentials:       getDBCredentials(),
		OidcAudience:        audience,
		OidcIssuer:          issuer,
		OidcAllowedAdGroups: allowedGroups.List,
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
