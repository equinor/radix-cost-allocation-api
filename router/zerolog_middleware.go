package router

import (
	"net/http"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/urfave/negroni/v3"
)

// Inspired by https://stackoverflow.com/a/50567022/2103434

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// NewZerologHandler injects and logs requests.
func NewZerologHandler(log zerolog.Logger) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		l := log.With().Logger()
		l.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("trace-id", xid.New().String())
		})
		start := time.Now()
		r = r.WithContext(l.WithContext(r.Context()))

		statusCodeWriter := newLoggingResponseWriter(w)
		next.ServeHTTP(statusCodeWriter, r)

		l.Info().
			Str("user-agent", r.Header.Get("User-Agent")).
			Str("remote-addr", r.RemoteAddr).
			Str("request", r.Method+" "+r.URL.Path).
			Str("query", r.URL.RawQuery).
			Dur("duration", time.Since(start)).
			Int("status", statusCodeWriter.statusCode).
			Msg(http.StatusText(statusCodeWriter.statusCode))
	}
}
