// httpsrv4
// Wrap a handler with middleware (a func(http.Handler) http.Handler).

package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func withHeader(value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-App", value)
			next.ServeHTTP(w, r)
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
