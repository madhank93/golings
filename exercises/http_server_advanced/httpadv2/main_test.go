// httpadv2
// Go 1.22's ServeMux supports method + wildcard patterns like
// "GET /users/{id}/posts/{postID}". r.PathValue("postID") returns that path
// segment — but ONLY if the name matches a {wildcard} in the pattern exactly.

package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}/posts/{postID}", func(w http.ResponseWriter, r *http.Request) {
		// FIXME: "post_id" is not a wildcard in the pattern, so PathValue returns
		// "". Read the value with the exact wildcard name: r.PathValue("postID").
		postID := r.PathValue("postID")
		if _, err := w.Write([]byte(postID)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	return mux
}

func TestPathValue(t *testing.T) {
	srv := httptest.NewServer(newRouter())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/users/7/posts/42")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "42" {
		t.Errorf("want 42, got %q", string(body))
	}
}
