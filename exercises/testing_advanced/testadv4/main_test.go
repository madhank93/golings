// testadv4 — benchmarks
// A benchmark measures performance; run it with `go test -bench=.`.

package main_test

import "testing"

func sumSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
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
