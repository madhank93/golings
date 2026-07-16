// pprof3
// Importing net/http/pprof for its side effects registers the /debug/pprof/
// handlers on http.DefaultServeMux. A server that uses its OWN mux won't expose
// them unless it forwards that path prefix to the default mux.

package main_test

import (
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"
)

func newDebugServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/debug/pprof/", http.DefaultServeMux)
	return mux
}

func TestPprofEndpoint(t *testing.T) {
	srv := httptest.NewServer(newDebugServer())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/debug/pprof/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want 200 from /debug/pprof/, got %d", resp.StatusCode)
	}
}
