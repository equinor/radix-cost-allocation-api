package controller

import (
	"net/http"

	"github.com/equinor/radix-cost-allocation-api/internal/accounts"
)

// Routes Holder of all routes
type Routes []Route

// Route Describe route
type Route struct {
	Path        string
	Method      string
	HandlerFunc RadixHandlerFunc
}

// RadixHandlerFunc Pattern for handler functions
type RadixHandlerFunc func(accounts.Accounts, http.ResponseWriter, *http.Request)
