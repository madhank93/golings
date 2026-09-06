package logingest

import (
	"errors"
	"strings"
	"testing"
	"time"
)

var parseAt = time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC)

func TestParseLineAcceptsWellFormedLines(t *testing.T) {
	cases := []struct {
		name string
		line string
		want Event
	}{
		{
			name: "plain",
			line: "api|info|request served",
			want: Event{Source: "api", Level: "info", Message: "request served", At: parseAt},
		},
		{
			name: "surrounding spaces are trimmed",
			line: "  api | warn |  disk filling up  ",
			want: Event{Source: "api", Level: "warn", Message: "disk filling up", At: parseAt},
		},
		{
			name: "separators inside the message are kept",
			line: "api|error|read failed|retrying|giving up",
			want: Event{Source: "api", Level: "error", Message: "read failed|retrying|giving up", At: parseAt},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseLine(tc.line, parseAt)
			if err != nil {
				t.Fatalf("ParseLine(%q) = error %v, want an event", tc.line, err)
			}
			if got != tc.want {
				t.Errorf("ParseLine(%q) =\n\t%+v\nwant\n\t%+v", tc.line, got, tc.want)
			}
		})
	}
}

func TestParseLineRejectsMalformedLines(t *testing.T) {
	cases := []struct {
		name string
		line string
	}{
		{"empty", ""},
		{"one field", "api"},
		{"two fields", "api|info"},
		{"empty source", "|info|hello"},
		{"empty message", "api|info|"},
		{"unknown level", "api|trace|hello"},
		{"only separators", "||"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseLine(tc.line, parseAt)
			if err == nil {
				t.Fatalf("ParseLine(%q) = nil error, want an error", tc.line)
			}
			if !errors.Is(err, ErrInvalidEvent) {
				t.Errorf("ParseLine(%q) error %v does not wrap ErrInvalidEvent", tc.line, err)
			}
		})
	}
}

// FuzzParseLine asserts two invariants rather than specific outputs, which is
// what makes it fuzzable: the fuzzer invents input, so there is no expected
// value to compare against — only properties that must hold for every input.
//
//  1. ParseLine never panics.
//  2. When it reports success, the event it returns really is valid.
//
// Run the seeds with `go test`, or go looking for counterexamples with
// `go test -fuzz=FuzzParseLine`.
func FuzzParseLine(f *testing.F) {
	f.Add("api|info|request served")
	f.Add("")
	f.Add("api")
	f.Add("api|info")
	f.Add("|||")
	f.Add("api|info|msg|with|pipes")
	f.Add("  api  |  debug  |  spaced  ")

	f.Fuzz(func(t *testing.T, line string) {
		// Invariant 1 holds implicitly: a panic here fails the fuzz target and
		// the offending input is written to testdata/fuzz for replay.
		e, err := ParseLine(line, parseAt)
		if err != nil {
			return
		}
		// Invariant 2: success must mean a genuinely valid event.
		if vErr := e.Validate(); vErr != nil {
			t.Errorf("ParseLine(%q) reported success but returned an invalid event: %v", line, vErr)
		}
		// A successful parse must not have invented or dropped content.
		if strings.TrimSpace(e.Source) != e.Source {
			t.Errorf("ParseLine(%q) left untrimmed whitespace in source %q", line, e.Source)
		}
	})
}

// BenchmarkParseLine uses b.Loop (Go 1.24), which replaces the older
// `for i := 0; i < b.N; i++` form. Run it with `go test -bench=.`
func BenchmarkParseLine(b *testing.B) {
	const line = "api|info|request served in 12ms"
	for b.Loop() {
		if _, err := ParseLine(line, parseAt); err != nil {
			b.Fatalf("ParseLine: %v", err)
		}
	}
}
