// interfaces4
// Interfaces can embed other interfaces. A type satisfies the combined
// interface by implementing every embedded method. Implement *Buffer.

// I AM NOT DONE
package main_test

import "testing"

type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

// ReadWriter embeds both Reader and Writer.
type ReadWriter interface {
	Reader
	Writer
}

type Buffer struct {
	data string
}

// FIXME: implement these on *Buffer so it satisfies ReadWriter:
//   func (b *Buffer) Read() string  { return b.data }
//   func (b *Buffer) Write(s string) { b.data += s }

func useRW(rw ReadWriter) string {
	rw.Write("hello")
	rw.Write(" world")
	return rw.Read()
}

func TestReadWriter(t *testing.T) {
	var rw ReadWriter = &Buffer{}
	if got := useRW(rw); got != "hello world" {
		t.Errorf("want %q, got %q", "hello world", got)
	}
}
