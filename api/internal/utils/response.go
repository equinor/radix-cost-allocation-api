package utils

import (
	"io"
	"net/http"

	radixhttp "github.com/equinor/radix-common/net/http"
	"github.com/rs/zerolog"
)

// ErrorResponseForServer Marshals error for server requester
func ErrorResponseForServer(w http.ResponseWriter, r *http.Request, apiError error) {
	err := radixhttp.ErrorResponseForServer(w, r, apiError)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msgf("%s %s: failed to write server response", r.Method, r.URL.Path)
	}
}

// JSONResponse Marshals response with header
func JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}) {
	err := radixhttp.JSONResponse(w, r, result)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msgf("%s %s: failed to write response", r.Method, r.URL.Path)
	}
}

// ReaderFileResponse writes the content from the reader to the response,
// and sets Content-Disposition=attachment; filename=<filename arg>
func ReaderFileResponse(w http.ResponseWriter, r *http.Request, reader io.Reader, fileName, contentType string) {
	err := radixhttp.ReaderFileResponse(w, reader, fileName, contentType)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msgf("%s %s: failed to write response", r.Method, r.URL.Path)
	}
}
