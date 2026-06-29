// stdlib1 — encoding/json
// Struct field tags control the JSON key names used during (un)marshaling.
// Add tags so the JSON keys map onto the struct fields.

// I AM NOT DONE
package main_test

import (
	"encoding/json"
	"testing"
)

// User should decode JSON like {"full_name":"alice","user_age":30}.
// (json matching is case-insensitive, so tags are required here because the
// keys differ by more than case.)
type User struct {
	// FIXME: add struct tags that map the JSON keys onto the fields.
	Name string
	Age  int
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
