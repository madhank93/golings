// httpsrv3
// Decode a JSON request body and encode a JSON response with encoding/json.

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

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var req echoReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(echoResp{Greeting: "Hello, " + req.Name + "!"})
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
