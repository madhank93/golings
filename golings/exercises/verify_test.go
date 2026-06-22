package exercises

import "testing"

func TestStatusString(t *testing.T) {
	cases := map[Status]string{
		StatusPending:  "Pending",
		StatusFailing:  "Failing",
		StatusLintFail: "LintFail",
		StatusDone:     "Done",
	}
	for s, want := range cases {
		if got := s.String(); got != want {
			t.Errorf("Status(%d).String() = %q, want %q", int(s), got, want)
		}
	}
}

// A pending exercise (marker present) must short-circuit to StatusPending
// without attempting to run, even when its path does not exist.
func TestVerifyPendingShortCircuits(t *testing.T) {
	// missing file => State() returns Pending (ReadFile error path)
	e := Exercise{Name: "ghost", Path: "does/not/exist/main.go", Mode: "compile"}
	status, _ := e.Verify()
	if status != StatusPending {
		t.Errorf("want StatusPending for missing/marked file, got %v", status)
	}
}
