// generics3
// Build a generic data structure: a Stack[T] that works for any element type.

package main_test

import "testing"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(x T) {
	s.items = append(s.items, x)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	x := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return x, true
}

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
