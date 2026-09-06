// binary2
// Pack and unpack a bitfield.

// I AM NOT DONE
package main_test

import "testing"

// The sixteen flag bits of a message header, most significant bit first:
//
//	 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
//	┌──┬────────────┬──┬──┬──┬──┬────────┬──────────┐
//	│QR│   opcode   │AA│TC│RD│RA│ zero   │  rcode   │
//	└──┴────────────┴──┴──┴──┴──┴────────┴──────────┘
//	 1       4        1  1  1  1     3         4      bits
//
// A field is a width and a shift. Reading one is "shift it down to bit 0, then
// mask off everything wider than the field"; writing one is the reverse. The
// mask is what stops a 5-bit value from corrupting its neighbour.
type Flags struct {
	Response    bool  // QR
	Opcode      uint8 // 4 bits
	Authorative bool  // AA
	Truncated   bool  // TC
	Recursion   bool  // RD
	Available   bool  // RA
	Rcode       uint8 // 4 bits
}

// ParseFlags unpacks the sixteen-bit flags word.
func ParseFlags(w uint16) Flags {
	// FIXME: shift each field down to bit 0 and mask it to its own width.
	// Response is done; the other six follow the same two steps.
	return Flags{
		Response: w>>15&1 == 1,
	}
}

// PackFlags is the inverse. The three reserved bits stay zero.
func PackFlags(f Flags) uint16 {
	var w uint16
	if f.Response {
		w |= 1 << 15
	}
	// FIXME: OR in the rest. Mask the multi-bit fields before shifting them:
	// an Opcode of 0xff must not spill into the bits either side of it.
	return w
}

func TestParseFlags(t *testing.T) {
	// 0x8180 = 1000 0001 1000 0000: a response, recursion desired and
	// available, no error.
	got := ParseFlags(0x8180)
	want := Flags{Response: true, Recursion: true, Available: true}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestParseFlagsWideFields(t *testing.T) {
	// 0x7803 = 0111 1000 0000 0011: opcode 15, rcode 3, nothing else set.
	got := ParseFlags(0x7803)
	if got.Opcode != 15 {
		t.Errorf("Opcode = %d, want 15", got.Opcode)
	}
	if got.Rcode != 3 {
		t.Errorf("Rcode = %d, want 3", got.Rcode)
	}
	if got.Response {
		t.Error("Response is set; bit 15 is clear in 0x7803")
	}
}

func TestPackFlagsRoundTrips(t *testing.T) {
	for _, w := range []uint16{0x0000, 0x8180, 0x7803, 0x0100, 0xf80f} {
		if got := PackFlags(ParseFlags(w)); got != w {
			t.Errorf("round trip of %#04x gave %#04x", w, got)
		}
	}
}

// A field wider than its slot must not bleed into the next one.
func TestPackFlagsMasksOverwideValues(t *testing.T) {
	got := PackFlags(Flags{Opcode: 0xff, Rcode: 0xff})
	if want := uint16(0x780f); got != want {
		t.Errorf("got %#04x, want %#04x — mask each field to its width", got, want)
	}
}
