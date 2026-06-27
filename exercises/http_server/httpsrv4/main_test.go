// httpsrv4
// Wrap a handler with middleware (a func(http.Handler) http.Handler).

// I AM NOT DONE
package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// withHeader returns middleware that sets the response header "X-App" before
// calling the next handler.
func withHeader(value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// FIXME: set w.Header().Set("X-App", value), then call
			// next.ServeHTTP(w, r) so the wrapped handler still runs.
		})
	}
}

func TestWithHeader(t *testing.T) {
	called := false
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})
	h := withHeader("golings")(final)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if !called {
		t.Error("middleware did not call the next handler")
	}
	if got := rec.Result().Header.Get("X-App"); got != "golings" {
		t.Errorf("want X-App golings, got %q", got)
	}
}
