package middleware

import (
	"net/http"
	"time"

	"github.com/equinor/radix-cost-allocation-api/internal/accounts"
	"github.com/equinor/radix-cost-allocation-api/internal/controller"
	radixhttp "github.com/equinor/radix-cost-allocation-api/internal/http"
	"github.com/rs/zerolog"
)

// RadixMiddleware The middleware between router and radix handler functions
type RadixMiddleware struct {
	Path    string
	Method  string
	next    controller.RadixHandlerFunc
	handled func(*RadixMiddleware, http.ResponseWriter, *http.Request, time.Time)
}

// NewRadixMiddleware Constructor for radix middleware
func NewRadixMiddleware(path, method string, next controller.RadixHandlerFunc, handled func(*RadixMiddleware, http.ResponseWriter, *http.Request, time.Time)) *RadixMiddleware {
	handler := &RadixMiddleware{
		Path:    path,
		Method:  method,
		next:    next,
		handled: handled,
	}
	return handler
}

// Handle Wraps radix handler methods
func (handler *RadixMiddleware) Handle(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	logger := zerolog.Ctx(r.Context())
	w.Header().Add("Access-Control-Allow-Origin", "*")

	defer func() {
		if handler.handled != nil {
			handler.handled(handler, w, r, startTime)
		}
	}()

	token, err := radixhttp.GetBearerTokenFromHeader(r)
	if err != nil {
		if err := radixhttp.ErrorResponse(w, r, err); err != nil {
			logger.Error().Err(err).Msg("unable to write auth error response")
		}
	}

	impersonation, err := radixhttp.GetImpersonationFromHeader(r)
	if err != nil {
		if err := radixhttp.ErrorResponse(w, r, radixhttp.UnexpectedError("Problems impersonating", err)); err != nil {
			logger.Error().Err(err).Msg("unable to write impersonating error response")
		}
	}

	accounts := accounts.NewAccounts(
		token,
		impersonation)

	handler.next(accounts, w, r)
}
