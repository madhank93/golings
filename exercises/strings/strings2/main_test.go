// strings2
// Count runes, not bytes, for correct Unicode lengths.

package main_test

import (
	"testing"
	"unicode/utf8"
)

// charCount returns the number of characters (runes) in s, not bytes.
func charCount(s string) int {
	return utf8.RuneCountInString(s)
}

func TestCharCount(t *testing.T) {
	if got := charCount("hello"); got != 5 {
		t.Errorf("ascii: want 5, got %d", got)
	}
	if got := charCount("héllo"); got != 5 {
		t.Errorf("accented: want 5, got %d", got)
	}
	if got := charCount("世界"); got != 2 {
		t.Errorf("cjk: want 2, got %d", got)
	}
}
