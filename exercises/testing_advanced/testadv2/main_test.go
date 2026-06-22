// testadv2 — fuzzing
// A fuzz test (func FuzzXxx(*testing.F)) feeds generated inputs to find edge
// cases; its seed corpus also runs during a normal `go test`.
// This reverse is buggy: it reverses BYTES, corrupting multibyte UTF-8.
// Fix it to reverse RUNES so the seed "Hello, 世界" stays valid.

// I AM NOT DONE
package main_test

import (
	"testing"
	"unicode/utf8"
)

// reverse should reverse s by runes so it round-trips and stays valid UTF-8.
func reverse(s string) string {
	// FIXME: reversing bytes breaks multibyte characters. Convert to []rune,
	// swap from both ends, and return string(runes).
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func FuzzReverse(f *testing.F) {
	for _, seed := range []string{"hello", "", "abc", "Hello, 世界"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, s string) {
		rev := reverse(s)
		if got := reverse(rev); got != s {
			t.Errorf("reverse twice = %q, want original %q", got, s)
		}
		if utf8.ValidString(s) && !utf8.ValidString(rev) {
			t.Errorf("reverse(%q) produced invalid UTF-8", s)
		}
	})
}
