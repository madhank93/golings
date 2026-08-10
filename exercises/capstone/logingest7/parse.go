// logingest7
// Stage seven is the text ingest path: agents that cannot speak JSON send
// pipe-separated lines instead. Parsing untrusted text is where panics live, so
// this is also where fuzzing earns its keep — the test file below asserts an
// invariant and lets the fuzzer go looking for input that breaks it.

package logingest

import (
	"fmt"
	"strings"
	"time"
)

// LineSeparator divides the fields of a plain-text log line.
const LineSeparator = "|"

// ParseLine parses a line of the form "source|level|message" into an Event
// stamped with at. It never panics.
func ParseLine(line string, at time.Time) (Event, error) {
	parts := strings.SplitN(line, LineSeparator, 3)
	if len(parts) != 3 {
		return Event{}, fmt.Errorf("%w: want %d fields separated by %q, got %d",
			ErrInvalidEvent, 3, LineSeparator, len(parts))
	}

	e := Event{
		Source:  strings.TrimSpace(parts[0]),
		Level:   strings.TrimSpace(parts[1]),
		Message: strings.TrimSpace(parts[2]),
		At:      at,
	}
	if err := e.Validate(); err != nil {
		return Event{}, err
	}
	return e, nil
}
