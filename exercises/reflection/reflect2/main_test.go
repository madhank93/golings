// reflect2
// Walk a struct's fields with reflection, calling a function for each string.

package main_test

import (
	"reflect"
	"testing"
)

func walk(x any, fn func(string)) {
	v := reflect.ValueOf(x)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			fn(v.Field(i).String())
		}
	}
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
