// typealias3
// A type ALIAS (`type MyByte = byte`, note the = sign) is not a new type — it is
// a second name for the SAME type. Values are interchangeable with no
// conversion, unlike a defined type.

package main_test

import "testing"

// FIXME: MyByte should be a true ALIAS of byte — write `type MyByte = byte`
// (with the =). As written it is a distinct defined type, so passing it to
// asByte below fails to compile.
type MyByte = byte

func asByte(b byte) byte { return b }

func TestAlias(t *testing.T) {
	var m MyByte = 65
	// Because MyByte is an alias of byte, it passes to asByte with NO conversion.
	if asByte(m) != 65 {
		t.Errorf("want 65, got %d", asByte(m))
	}
}
