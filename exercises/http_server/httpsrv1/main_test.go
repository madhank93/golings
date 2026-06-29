// httpsrv1
// Write an http.HandlerFunc and test it with httptest, no real network needed.

package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func TestGreetHandlerDefault(t *testing.T) {
	rec := httptest.NewRecorder()
	greetHandler(rec, httptest.NewRequest(http.MethodGet, "/greet", nil))
	if body, _ := io.ReadAll(rec.Result().Body); string(body) != "Hello, world!" {
		t.Errorf("want %q, got %q", "Hello, world!", string(body))
	}
}

func TestGreetHandlerWithName(t *testing.T) {
	rec := httptest.NewRecorder()
	greetHandler(rec, httptest.NewRequest(http.MethodGet, "/greet?name=Go", nil))
	if body, _ := io.ReadAll(rec.Result().Body); string(body) != "Hello, Go!" {
		t.Errorf("want %q, got %q", "Hello, Go!", string(body))
	}
}
