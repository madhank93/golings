// reflect2
// Walk a struct's fields with reflection, calling a function for each string.

// I AM NOT DONE
package main_test

import "testing"

// walk visits struct x and calls fn with the value of every string field.
func walk(x any, fn func(string)) {
	// FIXME: import "reflect". Take v := reflect.ValueOf(x); for each field
	// i in 0..v.NumField(), if v.Field(i).Kind() == reflect.String call
	// fn(v.Field(i).String()).
	_ = x
}

type person struct {
	Name string
	City string
	Age  int
}

func TestWalk(t *testing.T) {
	var got []string
	walk(person{Name: "Go", City: "Internet", Age: 15}, func(s string) {
		got = append(got, s)
	})
	want := []string{"Go", "Internet"}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("want %v, got %v", want, got)
	}
}
