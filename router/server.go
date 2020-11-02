package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/equinor/radix-cost-allocation-api/api/utils"
	"github.com/equinor/radix-cost-allocation-api/models"
	_ "github.com/equinor/radix-cost-allocation-api/swaggerui" // statik files
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
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
func NewServer(clusterName string, controllers ...models.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	tokenVerifier := getTokenVerifier(ctx)

	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	router.PathPrefix("/swaggerui/").Handler(sh)

	initializeAPIServer(router, controllers)

	initializeHealthEndpoint(router)

	serveMux := http.NewServeMux()
	serveMux.Handle(healthControllerPath, negroni.New(
		negroni.Wrap(router),
	))

	authenticationMiddleware := newAuthenticationMiddleware(tokenVerifier)
	authorizationMiddleware := newADGroupAuthorizationMiddleware(os.Getenv("AD_REPORT_READERS"), tokenVerifier)

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
		authenticationMiddleware,
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
		utils.NewRadixMiddleware(path, route.Method, route.HandlerFunc).Handle).Methods(route.Method)
}

func getTokenVerifier(ctx context.Context) *oidc.IDTokenVerifier {

	issuer := os.Getenv("TOKEN_ISSUER")

	provider, err := oidc.NewProvider(ctx, issuer)

	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		SkipClientIDCheck: true,
	}

	return provider.Verifier(oidcConfig)
}

func newAuthenticationMiddleware(verifier *oidc.IDTokenVerifier) negroni.HandlerFunc {
	ctx := context.Background()

	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := utils.GetBearerTokenFromHeader(r)

		if err != nil {
			w.WriteHeader(403)
			return
		}

		verified, err := verifier.Verify(ctx, token)

		if err != nil || verified == nil {
			w.WriteHeader(403)
			return
		}

		next(w, r)
	})
}

func newADGroupAuthorizationMiddleware(allowedADGroups string, verifier *oidc.IDTokenVerifier) negroni.HandlerFunc {
	ctx := context.Background()

	var allowedGroups struct {
		List []string `json:"groups"`
	}

	json.Unmarshal([]byte(allowedADGroups), &allowedGroups)

	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := utils.GetBearerTokenFromHeader(r)

		if err != nil {
			w.WriteHeader(403)
			return
		}

		var verified *oidc.IDToken
		verified, err = verifier.Verify(ctx, token)

		if err != nil || verified == nil {
			w.WriteHeader(403)
		}

		claims := &claims{}

		err = verified.Claims(claims)

		if err != nil {
			w.WriteHeader(403)
			return
		}

		for _, group := range claims.Groups {
			if find(allowedGroups.List, group) {
				next(w, r)
			}
		}

		w.WriteHeader(401)
		return

	})
}

type claims struct {
	Groups []string `json:"groups"`
	Email  string   `json:"email"`
}

func find(list []string, val string) bool {
	for _, item := range list {
		if strings.EqualFold(val, item) {
			return true
		}
	}

	return false
}
