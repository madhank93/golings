// pprof2
// runtime.ReadMemStats fills a MemStats snapshot of the heap. To measure how
// much memory a piece of code retains, read stats before and after — calling
// runtime.GC() first so previously-freed garbage is not counted.

// I AM NOT DONE
package main_test

import (
	"runtime"
	"testing"
)

// retainedBytes reports how many heap bytes alloc()'s result keeps alive.
func retainedBytes(alloc func() any) int64 {
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)

	keep := alloc()

	runtime.GC()
	runtime.ReadMemStats(&after)
	runtime.KeepAlive(keep)

	// FIXME: return the signed heap growth: int64(after.HeapAlloc) - int64(before.HeapAlloc).
	// Both fields are uint64, so subtracting them directly can wrap around to a huge
	// positive number. Right now it returns 0, so the test sees no allocation.
	return 0
}

func TestRetainedBytes(t *testing.T) {
	got := retainedBytes(func() any { return make([]byte, 1<<20) }) // 1 MiB

	const (
		want  = int64(1) << 20
		lower = want * 9 / 10
		upper = want * 2
	)
	// the delta is approximate because background garbage shifts the baseline
	// between the two snapshots, so an exact byte count would be flaky; the upper
	// bound also catches a delta that wrapped negative.
	if got < lower || got > upper {
		t.Errorf("got %d bytes, want in range [%d, %d]", got, lower, upper)
	}
}
