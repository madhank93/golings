// binary3
// Read length-prefixed frames off a stream.

package main_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"testing"
)

// MaxFrame is the largest payload ReadFrame will accept.
const MaxFrame = 1 << 20

// ErrFrameTooLarge is returned for a length prefix beyond MaxFrame.
var ErrFrameTooLarge = errors.New("frame too large")

// ReadFrame reads one length-prefixed message: a four-byte big-endian length,
// then exactly that many bytes of payload.
//
// The length is attacker-controlled, so it is checked before it is trusted.
// make([]byte, n) with an unchecked n from the wire is how a four-byte header
// becomes a four-gigabyte allocation.
//
// A clean end of stream — nothing at all left — is io.EOF and means the sender
// finished. A truncated frame is io.ErrUnexpectedEOF and means it did not.
func ReadFrame(r io.Reader) ([]byte, error) {
	var lenBuf [4]byte
	if _, err := io.ReadFull(r, lenBuf[:]); err != nil {
		return nil, err
	}
	n := binary.BigEndian.Uint32(lenBuf[:])
	if n > MaxFrame {
		return nil, fmt.Errorf("%w: %d bytes", ErrFrameTooLarge, n)
	}
	payload := make([]byte, n)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// WriteFrame writes one frame in the same format.
func WriteFrame(w io.Writer, payload []byte) error {
	var lenBuf [4]byte
	binary.BigEndian.PutUint32(lenBuf[:], uint32(len(payload)))
	if _, err := w.Write(lenBuf[:]); err != nil {
		return err
	}
	_, err := w.Write(payload)
	return err
}

func TestReadFrame(t *testing.T) {
	stream := []byte{0, 0, 0, 5, 'h', 'e', 'l', 'l', 'o'}
	got, err := ReadFrame(bytes.NewReader(stream))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "hello" {
		t.Errorf("got %q, want %q", got, "hello")
	}
}

// Framing exists so that several messages can share one stream: the reader
// must stop at the end of each payload, not read whatever is available.
func TestReadFrameManyFromOneStream(t *testing.T) {
	var buf bytes.Buffer
	for _, msg := range []string{"one", "", "three"} {
		if err := WriteFrame(&buf, []byte(msg)); err != nil {
			t.Fatal(err)
		}
	}

	for _, want := range []string{"one", "", "three"} {
		got, err := ReadFrame(&buf)
		if err != nil {
			t.Fatalf("reading %q: %v", want, err)
		}
		if string(got) != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}
	if _, err := ReadFrame(&buf); !errors.Is(err, io.EOF) {
		t.Errorf("after the last frame: %v, want io.EOF", err)
	}
}

func TestReadFrameTruncatedPayload(t *testing.T) {
	stream := []byte{0, 0, 0, 5, 'h', 'e'}
	_, err := ReadFrame(bytes.NewReader(stream))
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("err = %v, want io.ErrUnexpectedEOF", err)
	}
}

// The length says 4 GiB and the stream holds nothing. Allocating first would
// be a one-line denial of service.
func TestReadFrameRejectsAbsurdLength(t *testing.T) {
	stream := []byte{0xff, 0xff, 0xff, 0xff}
	_, err := ReadFrame(bytes.NewReader(stream))
	if !errors.Is(err, ErrFrameTooLarge) {
		t.Errorf("err = %v, want ErrFrameTooLarge", err)
	}
}
