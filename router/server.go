package router

import (
	"fmt"
	"github.com/equinor/radix-cost-allocation-api/models"
	_ "github.com/equinor/radix-cost-allocation-api/swaggerui" // statik files
	"github.com/equinor/radix-cost-allocation-api/utils"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"net/http"
	"os"
)

const (
	apiVersionRoute                 = "/api/v1"
	admissionControllerRootPath     = "/admissioncontrollers"
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
func NewServer(clusterName string, kubeUtil utils.KubeUtil, controllers ...models.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	router.PathPrefix("/swaggerui/").Handler(sh)

	initializeAPIServer(kubeUtil, router, controllers)

	initializeHealthEndpoint(router)

	serveMux := http.NewServeMux()
	serveMux.Handle(healthControllerPath, negroni.New(
		negroni.Wrap(router),
	))

	serveMux.Handle(fmt.Sprintf("%s/", admissionControllerRootPath), negroni.New(
		negroni.Wrap(router),
	))

	serveMux.Handle("/api/", negroni.New(
		negroni.HandlerFunc(utils.BearerTokenHeaderVerifyerMiddleware),
		negroni.Wrap(router),
	))

	// TODO: We should maybe have oauth to stop any non-radix user from beeing
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
			"http://localhost:3000",
			"http://localhost:3001",
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

func initializeAPIServer(kubeUtil utils.KubeUtil, router *mux.Router, controllers []models.Controller) {
	for _, controller := range controllers {
		for _, route := range controller.GetRoutes() {
			addHandlerRoute(kubeUtil, router, route)
		}
	}
}

func initializeHealthEndpoint(router *mux.Router) {
	router.HandleFunc(healthControllerPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
}

func addHandlerRoute(kubeUtil utils.KubeUtil, router *mux.Router, route models.Route) {
	path := apiVersionRoute + route.Path
	router.HandleFunc(path,
		utils.NewRadixMiddleware(kubeUtil, path, route.Method, route.HandlerFunc).Handle).Methods(route.Method)
}
