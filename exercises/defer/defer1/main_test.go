// defer1
// Deferred calls run when the surrounding function returns, in LIFO order
// (last deferred runs first). Using a named return value, the defers can even
// shape what is returned.

// I AM NOT DONE
package main_test

import "testing"

// order records the sequence in which deferred calls run. Defer appends of
// 1, then 2, then 3; because defers are LIFO, seq must end up [3, 2, 1].
func order() (seq []int) {
	// FIXME: add three deferred appends:
	//   defer func() { seq = append(seq, 1) }()
	//   defer func() { seq = append(seq, 2) }()
	//   defer func() { seq = append(seq, 3) }()
	return
}

func TestDeferOrder(t *testing.T) {
	got := order()
	want := []int{3, 2, 1}
	if len(got) != len(want) || got[0] != 3 || got[1] != 2 || got[2] != 1 {
		t.Errorf("want %v, got %v", want, got)
	}
}
