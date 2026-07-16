// httpadv1
// net/http's CrossOriginProtection (Go 1.25) blocks cross-site, state-changing
// requests using the browser's Sec-Fetch-Site metadata — CSRF defense with no
// tokens or cookies. You opt in by wrapping your handler with it.

package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// protect should wrap h so cross-site unsafe requests (like POST) are rejected.
func protect(h http.Handler) http.Handler {
	return http.NewCrossOriginProtection().Handler(h)
}

func TestCrossOriginBlocked(t *testing.T) {
	base := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	srv := httptest.NewServer(protect(base))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodPost, srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Sec-Fetch-Site", "cross-site")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Error("a cross-site POST should be rejected, but got 200")
	}
}
