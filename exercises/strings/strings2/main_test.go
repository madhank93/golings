// strings2
// A Go string is bytes, not characters. len(s) counts BYTES; one multibyte
// UTF-8 character (like é or 世) is several bytes. Count runes instead.

// I AM NOT DONE
package main_test

import "testing"

// charCount returns the number of characters (runes) in s, not bytes.
func charCount(s string) int {
	// FIXME: return utf8.RuneCountInString(s) (import "unicode/utf8").
	return len(s)
}

func TestCharCount(t *testing.T) {
	if got := charCount("hello"); got != 5 {
		t.Errorf("ascii: want 5, got %d", got)
	}
	if got := charCount("héllo"); got != 5 { // é is 2 bytes
		t.Errorf("accented: want 5, got %d", got)
	}
	if got := charCount("世界"); got != 2 { // each rune is 3 bytes
		t.Errorf("cjk: want 2, got %d", got)
	}
}
