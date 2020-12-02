package utils

import (
	"net/http"
	"time"

	"github.com/equinor/radix-cost-allocation-api/metrics"
	"github.com/equinor/radix-cost-allocation-api/models"
)

// RadixMiddleware The middleware between router and radix handler functions
type RadixMiddleware struct {
	path   string
	method string
	next   models.RadixHandlerFunc
}

// NewRadixMiddleware Constructor for radix middleware
func NewRadixMiddleware(path, method string, next models.RadixHandlerFunc) *RadixMiddleware {
	handler := &RadixMiddleware{
		path,
		method,
		next,
	}

	return handler
}

// Handle Wraps radix handler methods
func (handler *RadixMiddleware) Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	w.Header().Add("Access-Control-Allow-Origin", "*")

	defer func() {
		httpDuration := time.Since(start)
		metrics.AddRequestDuration(handler.path, handler.method, httpDuration)
	}()

	token, err := GetBearerTokenFromHeader(r)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	impersonation, err := getImpersonationFromHeader(r)
	if err != nil {
		ErrorResponse(w, r, UnexpectedError("Problems impersonating", err))
		return
	}

	accounts := models.NewAccounts(
		token,
		impersonation)

	handler.next(accounts, w, r)
}
