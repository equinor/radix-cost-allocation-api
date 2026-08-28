package http

import (
	"net/http"
	"strings"

	"github.com/equinor/radix-cost-allocation-api/internal/accounts"
	"github.com/pkg/errors"
)

// GetBearerTokenFromHeader gets bearer token from request header
func GetBearerTokenFromHeader(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("authorization")
	authArr := strings.Split(authorizationHeader, " ")
	var jwtToken string

	if len(authArr) != 2 {
		return "", errors.New("Authentication header is invalid: " + authorizationHeader)
	}

	jwtToken = authArr[1]
	return jwtToken, nil
}

// GetImpersonationFromHeader Gets Impersonation from request header
func GetImpersonationFromHeader(r *http.Request) (accounts.Impersonation, error) {
	impersonateUser := r.Header.Get("Impersonate-User")
	var impersonateGroups []string
	if impersonateGroupHeader := strings.TrimSpace(r.Header.Get("Impersonate-Group")); len(impersonateGroupHeader) > 0 {
		impersonateGroups = strings.Split(impersonateGroupHeader, ",")
		for i := range impersonateGroups {
			impersonateGroups[i] = strings.TrimSpace(impersonateGroups[i])
		}
	}

	return accounts.NewImpersonation(impersonateUser, impersonateGroups)
}
