package middleware

import (
	"net/http"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/urfave/negroni/v3"
)

// Inspired by https://stackoverflow.com/a/50567022/2103434

// NewZerologRequestLogger injects and logs requests.
func NewZerologRequestLogger(log zerolog.Logger) negroni.HandlerFunc {
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
			Dur("elapsed-ms", time.Since(start)).
			Int("status", statusCodeWriter.statusCode).
			Msg(http.StatusText(statusCodeWriter.statusCode))
	}
}

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
// Flush delegates to the underlying writer so streaming responses (SSE)
// still flush per chunk: the generated text/event-stream writer type-asserts
// http.Flusher on the outermost writer, and without this method the
// assertion fails and it falls back to a fully buffered io.Copy.
func (lrw *loggingResponseWriter) Flush() {
	if f, ok := lrw.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}
