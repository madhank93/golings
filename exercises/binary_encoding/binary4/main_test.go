// binary4
// Read a fixed-size page at an offset with io.ReaderAt.

// I AM NOT DONE
package main_test

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

// A paged file is an array of fixed-size pages addressed by number, the way
// every B-tree on disk is laid out:
//
//	0        100                pageSize            2*pageSize
//	├─────────┼─────────────────────┼───────────────────┤
//	│ 100-byte│  page 1 content     │  page 2           │  …
//	│  header │                     │                   │
//
// Pages are numbered from 1, so page n starts at (n-1)*pageSize. Page 1 is the
// exception: it shares its page with a 100-byte file header, so its content
// starts at offset 100.
const fileHeaderSize = 100

// ErrNoSuchPage is returned for a page number that is not in the file.
var ErrNoSuchPage = errors.New("no such page")

// ReadPage returns the content of page n from a paged file.
//
// io.ReaderAt reads at an absolute offset without a seek, so it is safe to call
// from several goroutines and needs no cursor of its own — which is why every
// random-access format reaches for it rather than Seek+Read.
//
// io.ReadFull turns a short read into an error: ReaderAt may legally return
// fewer bytes than asked for, and a page decoder handed a half page reports
// nonsense rather than an error.
func ReadPage(ra io.ReaderAt, pageSize, n int) ([]byte, error) {
	// FIXME: work out where page n starts and how long it is — remembering
	// that pages are numbered from 1 and that page 1 shares its page with the
	// 100-byte file header — then read exactly that many bytes at that offset.
	// A read that runs off the end of the file is ErrNoSuchPage, not an
	// io error to pass on.
	_ = io.NewSectionReader
	_ = errors.Is
	return nil, ErrNoSuchPage
}

// file builds a three-page file of pageSize 256: a 100-byte header, then
// page 1's content, then pages 2 and 3 filled with their own page number.
func file(pageSize int) []byte {
	buf := make([]byte, 0, pageSize*3)
	buf = append(buf, bytes.Repeat([]byte{0xaa}, fileHeaderSize)...)
	buf = append(buf, bytes.Repeat([]byte{1}, pageSize-fileHeaderSize)...)
	buf = append(buf, bytes.Repeat([]byte{2}, pageSize)...)
	buf = append(buf, bytes.Repeat([]byte{3}, pageSize)...)
	return buf
}

func TestReadPageSkipsTheFileHeader(t *testing.T) {
	const pageSize = 256
	got, err := ReadPage(bytes.NewReader(file(pageSize)), pageSize, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != pageSize-fileHeaderSize {
		t.Fatalf("page 1 is %d bytes, want %d", len(got), pageSize-fileHeaderSize)
	}
	if got[0] != 1 {
		t.Errorf("page 1 starts with %#02x — the 100-byte file header was not skipped", got[0])
	}
}

func TestReadPageIsOneBased(t *testing.T) {
	const pageSize = 256
	data := file(pageSize)

	for _, n := range []int{2, 3} {
		got, err := ReadPage(bytes.NewReader(data), pageSize, n)
		if err != nil {
			t.Fatalf("page %d: %v", n, err)
		}
		if len(got) != pageSize {
			t.Fatalf("page %d is %d bytes, want %d", n, len(got), pageSize)
		}
		if got[0] != byte(n) {
			t.Errorf("page %d starts with %d — off by one page?", n, got[0])
		}
	}
}

func TestReadPageOutOfRange(t *testing.T) {
	const pageSize = 256
	data := file(pageSize)

	for _, n := range []int{0, 4, 99} {
		if _, err := ReadPage(bytes.NewReader(data), pageSize, n); !errors.Is(err, ErrNoSuchPage) {
			t.Errorf("page %d: err = %v, want ErrNoSuchPage", n, err)
		}
	}
}
