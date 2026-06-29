// generics3
// Build a generic data structure: a Stack[T] that works for any element type.

// I AM NOT DONE
package main_test

import "testing"

// Stack is a last-in-first-out stack of any type T.
type Stack[T any] struct {
	items []T
}

// Push adds x to the top of the stack.
func (s *Stack[T]) Push(x T) {
	// FIXME: add x to the backing slice.
}

// Pop removes and returns the top item. The bool is false when the stack is
// empty (and the T is its zero value).
func (s *Stack[T]) Pop() (T, bool) {
	// FIXME: guard the empty case (zero value, false); otherwise remove and return the top item.
	var zero T
	return zero, false
}

// Len reports how many items are on the stack.
func (s *Stack[T]) Len() int { return len(s.items) }

func TestStackLIFO(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Fatalf("want len 3, got %d", s.Len())
	}
	for _, want := range []int{3, 2, 1} {
		got, ok := s.Pop()
		if !ok || got != want {
			t.Fatalf("want (%d, true), got (%d, %v)", want, got, ok)
		}
	}
	if _, ok := s.Pop(); ok {
		t.Error("popping an empty stack should return false")
	}
}

func TestStackString(t *testing.T) {
	var s Stack[string]
	s.Push("go")
	if got, ok := s.Pop(); !ok || got != "go" {
		t.Errorf("want (\"go\", true), got (%q, %v)", got, ok)
	}
}
