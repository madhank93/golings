// interfaces1
// An interface is a set of method signatures. A type satisfies it IMPLICITLY,
// just by having those methods — there is no "implements" keyword.
// Give Rectangle and Circle an Area method so both satisfy Shape.

// I AM NOT DONE
package main_test

import (
	"math"
	"testing"
)

// Shape is satisfied by anything with an Area() float64 method.
type Shape interface {
	Area() float64
}

type Rectangle struct{ W, H float64 }
type Circle struct{ R float64 }

// FIXME: give each type an Area method so it satisfies Shape.

func totalArea(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{Rectangle{W: 2, H: 3}, Circle{R: 1}}
	got := totalArea(shapes)
	want := 6.0 + math.Pi
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("want %v, got %v", want, got)
	}
}
