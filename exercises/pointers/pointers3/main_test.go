// pointers3
// A value parameter receives a COPY, so changes don't reach the caller.
// Only a pointer parameter can mutate the caller's data.

package main_test

import "testing"

type Counter struct {
	N int
}

// incValue takes a copy — incrementing here must NOT affect the caller.
func incValue(c Counter) {
	c.N++
}

// incPointer takes a pointer — incrementing here MUST affect the caller.
func incPointer(c *Counter) {
	c.N++
}

func TestValueVsPointer(t *testing.T) {
	c := Counter{N: 0}

	incValue(c)
	if c.N != 0 {
		t.Errorf("value parameter should not mutate caller; got %d", c.N)
	}

	incPointer(&c)
	if c.N != 1 {
		t.Errorf("pointer should mutate caller; want 1, got %d", c.N)
	}
}
