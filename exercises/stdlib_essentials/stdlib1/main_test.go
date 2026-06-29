// stdlib1 — encoding/json
// Struct field tags control the JSON key names used during (un)marshaling.

package main_test

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name string `json:"full_name"`
	Age  int    `json:"user_age"`
}

func parseUser(data []byte) (User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	return u, err
}

func TestParseUser(t *testing.T) {
	u, err := parseUser([]byte(`{"full_name":"alice","user_age":30}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Name != "alice" {
		t.Errorf("want Name alice, got %q", u.Name)
	}
	if u.Age != 30 {
		t.Errorf("want Age 30, got %d", u.Age)
	}
}
