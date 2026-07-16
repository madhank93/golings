// pprof1
// runtime/pprof writes a CPU profile you later inspect with `go tool pprof`.
// StartCPUProfile(w) begins sampling; the profile is only flushed to w when you
// call StopCPUProfile. Forget it and the output stays empty.

package main_test

import (
	"bytes"
	"io"
	"runtime/pprof"
	"testing"
)

// captureCPUProfile profiles work and writes the profile to w.
func captureCPUProfile(w io.Writer, work func()) error {
	if err := pprof.StartCPUProfile(w); err != nil {
		return err
	}
	defer pprof.StopCPUProfile()
	work()
	return nil
}

func TestCaptureCPUProfile(t *testing.T) {
	var buf bytes.Buffer
	err := captureCPUProfile(&buf, func() {
		sum := 0
		for i := range 1_000_000 {
			sum += i
		}
		_ = sum
	})
	if err != nil {
		t.Fatal(err)
	}
	if buf.Len() == 0 {
		t.Error("expected a non-empty CPU profile — did you call StopCPUProfile?")
	}
}
