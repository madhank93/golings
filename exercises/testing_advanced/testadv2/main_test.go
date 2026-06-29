// testadv2 — fuzzing
// A fuzz test feeds generated inputs to find edge cases. Reverse by runes so
// multibyte UTF-8 stays valid.

package main_test

import (
	"testing"
	"unicode/utf8"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
