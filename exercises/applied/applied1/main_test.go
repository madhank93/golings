// applied1 — sort.Interface
// Implementing sort.Interface (Len, Less, Swap) lets sort.Sort order ANY type.
// Len is provided; add Less and Swap.

// I AM NOT DONE
package main_test

import (
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// ByAge sorts a slice of Person by Age ascending.
type ByAge []Person

func (a ByAge) Len() int { return len(a) }

// FIXME: implement the other two methods of sort.Interface:
//   func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
//   func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func TestSortByAge(t *testing.T) {
	people := ByAge{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 35},
	}
	sort.Sort(people)

	if people[0].Name != "Bob" || people[1].Name != "Alice" || people[2].Name != "Carol" {
		t.Errorf("not sorted by age: %v", people)
	}
}
