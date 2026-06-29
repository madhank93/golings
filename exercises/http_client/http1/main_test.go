// http1 — HTTP client
// net/http makes requests. http.Get performs a GET; always close the body.

package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
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
