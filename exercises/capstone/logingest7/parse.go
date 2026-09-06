// logingest7
// Stage seven is the text ingest path: agents that cannot speak JSON send
// pipe-separated lines instead. Parsing untrusted text is where panics live, so
// this is also where fuzzing earns its keep — the test file below asserts an
// invariant and lets the fuzzer go looking for input that breaks it.

// I AM NOT DONE
package logingest

import "time"

// LineSeparator divides the fields of a plain-text log line.
const LineSeparator = "|"

// ParseLine parses a line of the form
//
//	source|level|message
//
// into an Event stamped with at. Surrounding spaces on each field are trimmed.
// The message is everything after the second separator, so it may itself
// contain separators.
//
// A line that does not have three fields, or that yields an event Validate
// rejects, is an error wrapping ErrInvalidEvent. ParseLine never panics,
// whatever it is handed.
func ParseLine(line string, at time.Time) (Event, error) {
	// FIXME: implement the parse.
	//
	// Use strings.SplitN(line, LineSeparator, 3) — not strings.Split. SplitN
	// with n=3 caps the result at three parts, which both keeps separators
	// inside the message and gives you a length you can check. Reach for
	// parts[2] after a plain Split on arbitrary input and you have written a
	// panic: "a|b" splits into two parts and the index is out of range.
	//
	// Then:
	//   - if len(parts) != 3, return an error wrapping ErrInvalidEvent,
	//   - build the Event from the three trimmed fields plus at,
	//   - run e.Validate() and return its error if the event is malformed,
	//   - otherwise return the event and nil.
	//
	// You will need to import "fmt" and "strings".
	//
	// Right now this accepts everything and returns the zero Event, which the
	// fuzz test catches immediately: it reports success for input that is not
	// a valid event.
	return Event{}, nil
}
