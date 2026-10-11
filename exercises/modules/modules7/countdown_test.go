package countdown_test

import (
	"slices"
	"testing"

	"example.com/countdown"
)

func TestFrom(t *testing.T) {
	want := []string{"3", "2", "1", "liftoff!"}
	if got := countdown.From(3); !slices.Equal(got, want) {
		t.Errorf("From(3) = %v, want %v", got, want)
	}
}
