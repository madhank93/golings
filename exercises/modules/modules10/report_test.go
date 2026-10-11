package report_test

import (
	"os/exec"
	"testing"

	"example.com/report"
)

func TestLine(t *testing.T) {
	if got, want := report.Line("apples", 12), "apples  12"; got != want {
		t.Errorf("Line = %q, want %q", got, want)
	}
}

// TestTidy fails while go.mod lists anything the imports don't need, or misses
// anything they do.
func TestTidy(t *testing.T) {
	out, err := exec.Command("go", "mod", "tidy", "-diff").CombinedOutput()
	if err != nil {
		t.Fatalf("go.mod does not match the imports:\n%s", out)
	}
}
