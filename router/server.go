package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/equinor/radix-cost-allocation-api/metrics"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/equinor/radix-common/models"
	radixnet "github.com/equinor/radix-common/net"
	radixhttp "github.com/equinor/radix-common/net/http"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"

	_ "github.com/equinor/radix-cost-allocation-api/swaggerui" // statik files
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

const (
	apiVersionRoute                 = "/api/v1"
	healthControllerPath            = "/health/"
	radixDNSZoneEnvironmentVariable = "RADIX_DNS_ZONE"
)

// Server Holds instance variables
type Server struct {
	Middleware  *negroni.Negroni
	clusterName string
	controllers []models.Controller
}

// NewServer Constructor function
func NewServer(clusterName string, authProvider auth.AuthProvider, controllers ...models.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	router.PathPrefix("/swaggerui/").Handler(sh)

	initializeAPIServer(router, controllers)

	initializeHealthEndpoint(router)

	serveMux := http.NewServeMux()
	serveMux.Handle(healthControllerPath, negroni.New(
		negroni.Wrap(router),
	))

	authenticationMiddleware := newAuthenticationMiddleware(authProvider)
	authorizationMiddleware := newADGroupAuthorizationMiddleware(os.Getenv("AD_REPORT_READERS"), authProvider)

	serveMux.Handle("/api/", negroni.New(
		authenticationMiddleware,
		negroni.Wrap(router),
	))

	serveMux.Handle("/api/v1/report", negroni.New(
		authenticationMiddleware,
		authorizationMiddleware,
		negroni.Wrap(router),
	))

	// TODO: We should maybe have oauth to stop any non-radix user from being
	// able to see the API
	serveMux.Handle("/swaggerui/", negroni.New(
		negroni.Wrap(router),
	))

	serveMux.Handle("/metrics", negroni.New(
		negroni.Wrap(promhttp.Handler()),
	))

	rec := negroni.NewRecovery()
	rec.PrintStack = false
	n := negroni.New(
		rec,
	)
	n.UseHandler(serveMux)

	server := &Server{
		n,
		clusterName,
		controllers,
	}

	return getCORSHandler(server)
}

func getCORSHandler(apiRouter *Server) http.Handler {
	radixDNSZone := os.Getenv(radixDNSZoneEnvironmentVariable)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3001",
			"http://localhost:3003",
			"http://localhost:8086", // For swaggerui testing
			// TODO: We should consider:
			// 1. "https://*.radix.equinor.com"
			// 2. Keep cors rules in ingresses
			fmt.Sprintf("https://console.%s", radixDNSZone),
			getHostName("web", "radix-web-console-qa", apiRouter.clusterName, radixDNSZone),
			getHostName("web", "radix-web-console-prod", apiRouter.clusterName, radixDNSZone),
			// Due to active-cluster
			getActiveClusterHostName("web", "radix-web-console-qa", radixDNSZone),
			getActiveClusterHostName("web", "radix-web-console-prod", radixDNSZone),
		},
		AllowCredentials: true,
		MaxAge:           600,
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "OPTIONS", "DELETE", "PATCH"},
	})
	return c.Handler(apiRouter.Middleware)
}

func getActiveClusterHostName(componentName, namespace, radixDNSZone string) string {
	return fmt.Sprintf("https://%s-%s.%s", componentName, namespace, radixDNSZone)
}

func getHostName(componentName, namespace, clustername, radixDNSZone string) string {
	return fmt.Sprintf("https://%s-%s.%s.%s", componentName, namespace, clustername, radixDNSZone)
}

func initializeAPIServer(router *mux.Router, controllers []models.Controller) {
	for _, controller := range controllers {
		for _, route := range controller.GetRoutes() {
			addHandlerRoute(router, route)
		}
	}
}

func initializeHealthEndpoint(router *mux.Router) {
	router.HandleFunc(healthControllerPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
}

func addHandlerRoute(router *mux.Router, route models.Route) {
	path := apiVersionRoute + route.Path
	router.HandleFunc(path,
		radixnet.NewRadixMiddleware(path, route.Method, route.HandlerFunc,
			func(handler *radixnet.RadixMiddleware, w http.ResponseWriter, r *http.Request, started time.Time) {
				httpDuration := time.Since(started)
				metrics.AddRequestDuration(handler.path, handler.method, httpDuration)
			}).Handle).
		Methods(route.Method)
}

func newAuthenticationMiddleware(authProvider auth.AuthProvider) negroni.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := radixhttp.GetBearerTokenFromHeader(r)

		if err != nil {
			log.Info("Could not get token from header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		verified, err := authProvider.VerifyToken(ctx, token)

		if err != nil || verified == nil {
			log.Debugf("Could not verify token. Error: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func newADGroupAuthorizationMiddleware(allowedADGroups string, authProvider auth.AuthProvider) negroni.HandlerFunc {
	ctx := context.Background()

	var allowedGroups struct {
		List []string `json:"groups"`
	}

	err := json.Unmarshal([]byte(allowedADGroups), &allowedGroups)
	if err != nil {
		log.Errorf("could not parse json for allowedADGroups")
	}

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := radixhttp.GetBearerTokenFromHeader(r)

		if err != nil {
			log.Info("Could not get token from header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var verified auth.IDToken
		verified, err = authProvider.VerifyToken(ctx, token)

		if err != nil || verified == nil {
			log.Debugf("Unable to verify token. Error: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
		}

		claims := &auth.Claims{}

		err = verified.GetClaims(claims)

		if err != nil {
			log.Debugf("Could not get claims from token. Error: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, group := range claims.Groups {
			if find(allowedGroups.List, group) {
				next(w, r)
			}
		}

		log.Debugf("User does not have correct AD group access. Logged AD groups for user: %v ", claims.Groups)
		w.WriteHeader(http.StatusForbidden)
		return

	}
}

func find(list []string, val string) bool {
	for _, item := range list {
		if strings.EqualFold(val, item) {
			return true
		}
	}

	return false
}
