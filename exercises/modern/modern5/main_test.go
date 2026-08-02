// modern5
// Go 1.22 changed for-loop semantics: each iteration declares NEW loop
// variables instead of reusing one across the whole loop. Closures and pointers
// created inside the loop therefore capture a value of their own — the classic
// "all my goroutines saw the last item" bug is gone, but only when the variable
// really is declared by the loop.

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
)

// labelers returns one closure per item, each reporting its own item.
func labelers(items []string) []func() string {
	var fns []func() string
	// FIXME: `item` is hoisted out of the loop and merely assigned by `=`, so
	// there is a single variable that every closure shares — the pre-1.22
	// shape. Let the loop declare it with `:=` so each iteration gets its own.
	var item string
	for _, item = range items {
		fns = append(fns, func() string { return item })
	}
	return fns
}

func TestLabelers(t *testing.T) {
	var got []string
	for _, fn := range labelers([]string{"alpha", "beta", "gamma"}) {
		got = append(got, fn())
	}
	want := []string{"alpha", "beta", "gamma"}
	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
