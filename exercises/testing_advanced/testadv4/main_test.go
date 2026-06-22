// testadv4 — benchmarks
// A benchmark (func BenchmarkXxx(*testing.B)) measures performance; run it with
// `go test -bench=.`. The Go 1.24+ form is `for b.Loop() { ... }`.
// The benchmark below is complete — your job is to implement sumSlice so the
// correctness test passes (normal `go test` runs the Test, not the Benchmark).

// I AM NOT DONE
package main_test

import "testing"

// sumSlice returns the sum of all elements.
func sumSlice(nums []int) int {
	// FIXME: range over nums and add each value.
	return 0
}

func TestSumSlice(t *testing.T) {
	if got := sumSlice([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("want 10, got %d", got)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	for b.Loop() {
		sumSlice(data)
	}
}
