// binary1
// Decode a fixed-size big-endian header.

package main_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"
)

// HeaderSize is the wire size of a Header: six 16-bit fields.
const HeaderSize = 12

// Header is a twelve-byte message header. Every field is a uint16 in network
// byte order — big-endian, most significant byte first — which is what "network
// byte order" means and why the struct's Go field order says nothing about the
// bytes.
type Header struct {
	ID        uint16
	Flags     uint16
	Questions uint16
	Answers   uint16
	Authority uint16
	Extra     uint16
}

// DecodeHeader reads exactly HeaderSize bytes from r and decodes them.
//
// A stream that ends early is io.ErrUnexpectedEOF, not io.EOF: EOF means "there
// was nothing", and a half-header means "there was something and it was
// truncated". io.ReadFull draws that distinction for you.
func DecodeHeader(r io.Reader) (Header, error) {
	var buf [HeaderSize]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return Header{}, err
	}
	return Header{
		ID:        binary.BigEndian.Uint16(buf[0:2]),
		Flags:     binary.BigEndian.Uint16(buf[2:4]),
		Questions: binary.BigEndian.Uint16(buf[4:6]),
		Answers:   binary.BigEndian.Uint16(buf[6:8]),
		Authority: binary.BigEndian.Uint16(buf[8:10]),
		Extra:     binary.BigEndian.Uint16(buf[10:12]),
	}, nil
}

// EncodeHeader is the inverse: HeaderSize bytes, same order, same endianness.
func EncodeHeader(h Header) []byte {
	buf := make([]byte, HeaderSize)
	binary.BigEndian.PutUint16(buf[0:2], h.ID)
	binary.BigEndian.PutUint16(buf[2:4], h.Flags)
	binary.BigEndian.PutUint16(buf[4:6], h.Questions)
	binary.BigEndian.PutUint16(buf[6:8], h.Answers)
	binary.BigEndian.PutUint16(buf[8:10], h.Authority)
	binary.BigEndian.PutUint16(buf[10:12], h.Extra)
	return buf
}

var wire = []byte{
	0x04, 0xd2, // ID        = 1234
	0x81, 0x80, // Flags     = 0x8180
	0x00, 0x01, // Questions = 1
	0x00, 0x02, // Answers   = 2
	0x00, 0x00, // Authority = 0
	0x00, 0x00, // Extra     = 0
}

func TestDecodeHeader(t *testing.T) {
	got, err := DecodeHeader(bytes.NewReader(wire))
	if err != nil {
		t.Fatal(err)
	}
	want := Header{ID: 1234, Flags: 0x8180, Questions: 1, Answers: 2}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// Little-endian would read 0x04d2 as 0xd204 = 53764. This is the whole bug.
func TestDecodeHeaderIsBigEndian(t *testing.T) {
	h, err := DecodeHeader(bytes.NewReader(wire))
	if err != nil {
		t.Fatal(err)
	}
	if h.ID == 0xd204 {
		t.Fatal("the bytes were read little-endian; the wire is big-endian")
	}
	if h.ID != 0x04d2 {
		t.Errorf("ID = %#04x, want 0x04d2", h.ID)
	}
}

func TestDecodeHeaderShortInput(t *testing.T) {
	_, err := DecodeHeader(bytes.NewReader(wire[:5]))
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("err = %v, want io.ErrUnexpectedEOF", err)
	}
}

func TestHeaderRoundTrips(t *testing.T) {
	h := Header{ID: 0xbeef, Flags: 0x0100, Questions: 1, Extra: 7}
	got, err := DecodeHeader(bytes.NewReader(EncodeHeader(h)))
	if err != nil {
		t.Fatal(err)
	}
	if got != h {
		t.Errorf("round trip: got %+v, want %+v", got, h)
	}
	if n := len(EncodeHeader(h)); n != HeaderSize {
		t.Errorf("encoded %d bytes, want %d", n, HeaderSize)
	}
}
