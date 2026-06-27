// reflect1
// Inspect a value's kind at runtime with the reflect package.

// I AM NOT DONE
package main_test

import "testing"

// describe returns the kind of x as a string — e.g. 42 -> "int",
// "hi" -> "string", []int{} -> "slice".
func describe(x any) string {
	// FIXME: import "reflect" and return reflect.TypeOf(x).Kind().String().
	return ""
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{42, "int"},
		{"hi", "string"},
		{3.14, "float64"},
		{[]int{1}, "slice"},
		{map[string]int{}, "map"},
	}
	for _, c := range cases {
		if got := describe(c.in); got != c.want {
			t.Errorf("describe(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
