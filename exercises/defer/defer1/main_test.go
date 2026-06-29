// defer1
// Deferred calls run when the surrounding function returns, in LIFO order.

package main_test

import "testing"

func order() (seq []int) {
	defer func() { seq = append(seq, 1) }()
	defer func() { seq = append(seq, 2) }()
	defer func() { seq = append(seq, 3) }()
	return
}

func TestDeferOrder(t *testing.T) {
	got := order()
	want := []int{3, 2, 1}
	if len(got) != len(want) || got[0] != 3 || got[1] != 2 || got[2] != 1 {
		t.Errorf("want %v, got %v", want, got)
	}
}
