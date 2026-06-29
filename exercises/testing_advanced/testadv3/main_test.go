// testadv3 — httptest
// net/http/httptest exercises HTTP handlers without a real network.

package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()

	pingHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("want status 200, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if string(body) != "pong" {
		t.Errorf("want body %q, got %q", "pong", string(body))
	}
}
