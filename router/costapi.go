package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/equinor/radix-common/models"
	radixnet "github.com/equinor/radix-common/net"
	radixhttp "github.com/equinor/radix-common/net/http"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	"github.com/equinor/radix-cost-allocation-api/metrics"
	"github.com/equinor/radix-cost-allocation-api/swaggerui"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni/v3"
)

const (
	apiVersionRoute                 = "/api/v1"
	healthControllerPath            = "/health/"
	radixDNSZoneEnvironmentVariable = "RADIX_DNS_ZONE"
	swaggerUIPath                   = "/swaggerui"
)

// NewHandler Constructor function
func NewHandler(clusterName string, allowedAdGroups []string, authProvider auth.AuthProvider, controllers ...models.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	initializeSwaggerUI(router)
	initializeAPIServer(router, controllers)
	initializeHealthEndpoint(router)

	serveMux := http.NewServeMux()
	serveMux.Handle(healthControllerPath, negroni.New(
		negroni.Wrap(router),
	))

	authenticationMiddleware := newAuthenticationMiddleware(authProvider)
	authorizationMiddleware := newADGroupAuthorizationMiddleware(allowedAdGroups, authProvider)

	serveMux.Handle("/api/", negroni.New(
		authenticationMiddleware,
		negroni.Wrap(router),
	))

	serveMux.Handle("/api/v1/report", negroni.New(
		authorizationMiddleware,
		negroni.Wrap(router),
	))

	// TODO: We should maybe have oauth to stop any non-radix user from being
	// able to see the API
	serveMux.Handle("/swaggerui/", negroni.New(
		negroni.Wrap(router),
	))

	rec := negroni.NewRecovery()
	rec.PrintStack = false
	n := negroni.New(
		rec,
		NewZerologHandler(log.Logger),
	)

	n.UseHandler(serveMux)

	return applyCORS(clusterName, n)
}

func applyCORS(clusterName string, handler http.Handler) http.Handler {
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
			getHostName("web", "radix-web-console-qa", clusterName, radixDNSZone),
			getHostName("web", "radix-web-console-prod", clusterName, radixDNSZone),
			// Due to active-cluster
			getActiveClusterHostName("web", "radix-web-console-qa", radixDNSZone),
			getActiveClusterHostName("web", "radix-web-console-prod", radixDNSZone),
		},
		AllowCredentials: true,
		MaxAge:           600,
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "OPTIONS", "DELETE", "PATCH"},
	})
	return c.Handler(handler)
}

func getActiveClusterHostName(componentName, namespace, radixDNSZone string) string {
	return fmt.Sprintf("https://%s-%s.%s", componentName, namespace, radixDNSZone)
}

func getHostName(componentName, namespace, clustername, radixDNSZone string) string {
	return fmt.Sprintf("https://%s-%s.%s.%s", componentName, namespace, clustername, radixDNSZone)
}

func initializeSwaggerUI(router *mux.Router) {
	swaggerFsHandler := http.FileServer(http.FS(swaggerui.FS()))
	swaggerui := http.StripPrefix(swaggerUIPath, swaggerFsHandler)
	router.PathPrefix(swaggerUIPath).Handler(swaggerui)
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
				metrics.AddRequestDuration(handler.Path, handler.Method, httpDuration)
			}).Handle).
		Methods(route.Method)
}

func newAuthenticationMiddleware(authProvider auth.AuthProvider) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := radixhttp.GetBearerTokenFromHeader(r)

		if err != nil {
			zerolog.Ctx(r.Context()).Info().Msg("Could not get token from header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		verified, err := authProvider.VerifyToken(r.Context(), token)
		// TODO: Validate token audience
		if err != nil || verified == nil {
			zerolog.Ctx(r.Context()).Debug().Err(err).Msg("Could not verify token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func newADGroupAuthorizationMiddleware(allowedADGroups []string, authProvider auth.AuthProvider) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := radixhttp.GetBearerTokenFromHeader(r)

		if err != nil {
			zerolog.Ctx(r.Context()).Info().Msg("Could not get token from header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var verified auth.IDToken
		verified, err = authProvider.VerifyToken(r.Context(), token)

		if err != nil || verified == nil {
			zerolog.Ctx(r.Context()).Debug().Err(err).Msg("Unable to verify token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &auth.Claims{}

		err = verified.GetClaims(claims)

		if err != nil {
			zerolog.Ctx(r.Context()).Debug().Err(err).Msg("Could not get claims from token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, group := range claims.Groups {
			if find(allowedADGroups, group) {
				next(w, r)
				return
			}
		}

		zerolog.Ctx(r.Context()).Debug().Strs("ad-groups", claims.Groups).Msgf("User does not have correct AD group access")
		w.WriteHeader(http.StatusForbidden)
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
