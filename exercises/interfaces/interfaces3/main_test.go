// interfaces3
// Recover the concrete type held by an interface value with a type switch
// (or a comma-ok type assertion).

// I AM NOT DONE
package main_test

import (
	"testing"
)

// describe reports v as "int: N" for ints, "string: S" for strings,
// or "unknown" for anything else.
func describe(v any) string {
	// FIXME: use a type switch on v.(type) (import "fmt"):
	//   case int:    return fmt.Sprintf("int: %d", x)
	//   case string: return fmt.Sprintf("string: %s", x)
	//   default:     return "unknown"
	return ""
}

func TestDescribe(t *testing.T) {
	if got := describe(42); got != "int: 42" {
		t.Errorf("int case: got %q", got)
	}
	if got := describe("hi"); got != "string: hi" {
		t.Errorf("string case: got %q", got)
	}
	if got := describe(3.14); got != "unknown" {
		t.Errorf("default case: got %q", got)
	}
}
