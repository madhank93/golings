// maps2
// Initialize a map with a literal.
//
// A map literal builds and fills a map in one expression: map[K]V{k: v, ...}.

// I AM NOT DONE
package main_test

import "testing"

// FIXME: return a map[string]int literal with John=30 and Ana=21.
func ages() map[string]int {
	return map{}
}

func TestAges(t *testing.T) {
	m := ages()
	if m["John"] != 30 || m["Ana"] != 21 {
		t.Errorf("want John=30 Ana=21, got John=%d Ana=%d", m["John"], m["Ana"])
	}
}
