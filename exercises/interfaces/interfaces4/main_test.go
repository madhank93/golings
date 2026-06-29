// interfaces4
// Interfaces can embed other interfaces. A type satisfies the combined
// interface by implementing every embedded method.

package main_test

import "testing"

type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

type ReadWriter interface {
	Reader
	Writer
}

type Buffer struct {
	data string
}

func (b *Buffer) Read() string   { return b.data }
func (b *Buffer) Write(s string) { b.data += s }

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
