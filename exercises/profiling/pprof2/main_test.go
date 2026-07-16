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
func retainedBytes(alloc func() any) uint64 {
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)

	keep := alloc()

	runtime.GC()
	runtime.ReadMemStats(&after)
	runtime.KeepAlive(keep)

	// FIXME: return the heap growth: after.HeapAlloc - before.HeapAlloc. Right
	// now it returns 0, so the test sees no allocation.
	return 0
}

func TestRetainedBytes(t *testing.T) {
	got := retainedBytes(func() any { return make([]byte, 1<<20) }) // 1 MiB
	if got < 1<<20 {
		t.Errorf("want at least 1 MiB retained, got %d bytes", got)
	}
}
