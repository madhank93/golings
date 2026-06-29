// concpat3 — pipeline
// A pipeline chains stages connected by channels: each stage reads from the
// previous stage and emits to the next. Implement the `double` stage.

// I AM NOT DONE
package main_test

import (
	"testing"
)

// generate emits nums on a channel, then closes it.
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// double reads from in and emits each value times two, closing out when in
// is exhausted.
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// FIXME: range the input, send each doubled value, then close out after the loop.
		close(out)
	}()
	return out
}

func TestPipeline(t *testing.T) {
	var got []int
	for v := range double(generate(1, 2, 3)) {
		got = append(got, v)
	}

	want := []int{2, 4, 6}
	if len(got) != len(want) {
		t.Fatalf("want %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("at %d: want %d, got %d", i, want[i], got[i])
		}
	}
}
