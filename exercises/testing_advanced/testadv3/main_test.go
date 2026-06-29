// testadv3 — httptest
// net/http/httptest exercises HTTP handlers without a real network: build a
// request, record the response, assert on it. Implement the handler.

// I AM NOT DONE
package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// pingHandler should respond with status 200 and the body "pong".
func pingHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: write the body "pong" to w and check the write error.
	_ = w
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
