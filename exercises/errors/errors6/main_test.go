// errors6
// errors.Is walks the ENTIRE wrap chain: fmt.Errorf("...: %w", err) links a new
// message onto an existing error, and errors.Is finds a sentinel no matter how
// many layers deep it sits. Using %v instead of %w flattens the error to plain
// text and severs the chain.

package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrNotFound = errors.New("not found")

// openConfig wraps ErrNotFound three layers deep. Every layer must keep the
// original reachable via errors.Is.
func openConfig() error {
	level1 := fmt.Errorf("read file: %w", ErrNotFound)
	// FIXME: the middle layer uses %v, which flattens level1 to text and breaks
	// the chain — errors.Is can no longer reach ErrNotFound. Wrap with %w.
	level2 := fmt.Errorf("load config: %w", level1)
	level3 := fmt.Errorf("startup: %w", level2)
	return level3
}

func TestOpenConfig(t *testing.T) {
	err := openConfig()
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("errors.Is should reach ErrNotFound through the wrap chain; got %v", err)
	}
}
