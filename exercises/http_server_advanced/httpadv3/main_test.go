// httpadv3
// httputil.ReverseProxy.Rewrite is the modern hook (replacing the deprecated
// Director field) for retargeting a request. Inside it, pr.SetURL(target) points
// the outbound request at the backend while preserving the inbound path.

// I AM NOT DONE
package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func newProxy(target *url.URL) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			// FIXME: the outbound request is never pointed at the backend, so it
			// forwards nowhere. Call pr.SetURL(target) here.
			_ = pr
		},
	}
}

func TestProxyForwards(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("from backend")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
	defer backend.Close()

	target, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatal(err)
	}
	front := httptest.NewServer(newProxy(target))
	defer front.Close()

	resp, err := http.Get(front.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "from backend" {
		t.Errorf("want %q, got %q", "from backend", string(body))
	}
}
