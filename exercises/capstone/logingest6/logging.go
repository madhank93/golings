// logingest6
// Stage six is observability. A log line built with fmt.Sprintf is fine to read
// and miserable to query; log/slog emits key/value pairs that a log aggregator
// can filter on. The request logger is middleware, so it wraps any handler
// without the handler knowing.

package logingest

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

// NewLogger returns a logger writing JSON records to w, discarding anything
// below level.
func NewLogger(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level}))
}

// statusRecorder wraps an http.ResponseWriter to remember the status code that
// was written.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

// WriteHeader records the code before passing it on.
func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// WithLogging returns a handler that runs next and logs one record per request.
func WithLogging(next http.Handler, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// 200 is the status a handler sends by writing a body without ever
		// calling WriteHeader, so it is the right default here.
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		log.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"duration", time.Since(start),
		)
	})
}
