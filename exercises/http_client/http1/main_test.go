// http1 — HTTP client
// net/http makes requests. http.Get performs a GET; always close the response
// body. The test spins up a local httptest.Server so the call is deterministic.

// I AM NOT DONE
package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// fetch performs a GET request to url and returns the response body as a string.
func fetch(url string) (string, error) {
	// FIXME: GET the url, defer-close the body, and read it all.
	return "", nil
}

func TestFetch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "pong")
	}))
	defer srv.Close()

	got, err := fetch(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "pong" {
		t.Errorf("want pong, got %q", got)
	}
}
