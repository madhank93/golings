// logingest6
// Stage six is observability. A log line built with fmt.Sprintf is fine to read
// and miserable to query; log/slog emits key/value pairs that a log aggregator
// can filter on. The request logger is middleware, so it wraps any handler
// without the handler knowing.

// I AM NOT DONE
package logingest

import (
	"io"
	"log/slog"
	"net/http"
)

// NewLogger returns a logger writing JSON records to w, discarding anything
// below level.
func NewLogger(w io.Writer, level slog.Level) *slog.Logger {
	// FIXME: build and return a JSON logger:
	//
	//	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level}))
	//
	// Right now this ignores both w and level and returns the package default,
	// which writes text to stderr — so the tests capture nothing.
	return slog.Default()
}

// statusRecorder wraps an http.ResponseWriter to remember the status code that
// was written. The ResponseWriter interface offers no way to read it back, and
// middleware that wants to log the status has to observe it on the way out.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

// WriteHeader records the code before passing it on.
func (r *statusRecorder) WriteHeader(code int) {
	// FIXME: remember code in r.status, then call the embedded
	// ResponseWriter's WriteHeader so the response is actually sent.
	//
	// Forgetting the second half is the classic middleware bug: the status is
	// logged correctly and never reaches the client.
	r.ResponseWriter.WriteHeader(code)
}

// WithLogging returns a handler that runs next and logs one record per request
// with the method, path, status and duration.
func WithLogging(next http.Handler, log *slog.Logger) http.Handler {
	// FIXME: return an http.HandlerFunc that
	//
	//  1. notes the start time,
	//  2. wraps w in a *statusRecorder whose status starts at
	//     http.StatusOK — a handler that writes a body without ever calling
	//     WriteHeader still sends 200, so 200 is the correct default and 0 is
	//     not a status code,
	//  3. calls next.ServeHTTP with the recorder,
	//  4. logs at Info with the attributes "method", "path", "status" and
	//     "duration", using the message "request".
	//
	// You will need to import "time".
	//
	// Right now next is returned unwrapped, so nothing is ever logged.
	return next
}
