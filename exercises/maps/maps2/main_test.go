// maps2
// Initialize a map with a literal.
//
// A map literal builds and fills a map in one expression: map[K]V{k: v, ...}.

package main_test

import "testing"

func ages() map[string]int {
	return map[string]int{"John": 30, "Ana": 21}
}

func TestAges(t *testing.T) {
	m := ages()
	if m["John"] != 30 || m["Ana"] != 21 {
		t.Errorf("want John=30 Ana=21, got John=%d Ana=%d", m["John"], m["Ana"])
	}
}
