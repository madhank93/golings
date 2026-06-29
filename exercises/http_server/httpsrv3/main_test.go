// httpsrv3
// Decode a JSON request body and encode a JSON response with encoding/json.

// I AM NOT DONE
package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type echoReq struct {
	Name string `json:"name"`
}

type echoResp struct {
	Greeting string `json:"greeting"`
}

// echoHandler decodes {"name": "..."} and responds with
// {"greeting": "Hello, <name>!"} and Content-Type application/json.
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: decode the JSON body, set the JSON content-type, then encode the reply.
	_ = w
	_ = r
}

func TestEchoHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(`{"name":"Go"}`))
	rec := httptest.NewRecorder()
	echoHandler(rec, req)

	res := rec.Result()
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("want Content-Type application/json, got %q", ct)
	}
	var got echoResp
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Greeting != "Hello, Go!" {
		t.Errorf("want greeting %q, got %q", "Hello, Go!", got.Greeting)
	}
}
