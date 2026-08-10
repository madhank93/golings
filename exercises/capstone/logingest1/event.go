// logingest1
// The capstone builds one program across eight stages: a concurrent HTTP
// log-ingest service. Stage one is the data model — the Event that every later
// stage validates, stores, serves, batches, logs and profiles.
//
// This is the first exercise that spans more than one file. The implementation
// lives here, the tests live next door, and golings compiles the whole
// directory rather than a single file.

// Carried forward from logingest1, already solved. Read it, do not edit it —
// the stage you are working on is store.go.

package logingest

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

// ErrInvalidEvent is the sentinel every validation failure wraps. Callers care
// that an event was rejected, not which field was wrong, so they match on this
// one value with errors.Is while the wrapped text still names the field.
var ErrInvalidEvent = errors.New("invalid event")

// Levels is the closed set of severities the service accepts.
var Levels = []string{"debug", "info", "warn", "error"}

// Event is a single log line as it arrives on the wire. The struct tags are
// what the stage-three JSON handler will decode into.
type Event struct {
	Source  string    `json:"source"`
	Level   string    `json:"level"`
	Message string    `json:"message"`
	At      time.Time `json:"at"`
}

// Validate reports whether e is well-formed, returning an error that wraps
// ErrInvalidEvent when it is not.
//
// An event is valid when Source is non-empty, Level is one of Levels, Message
// is non-empty, and At is not the zero time.
func (e Event) Validate() error {
	switch {
	case e.Source == "":
		return fmt.Errorf("%w: source is empty", ErrInvalidEvent)
	case !slices.Contains(Levels, e.Level):
		return fmt.Errorf("%w: level %q is not one of %v", ErrInvalidEvent, e.Level, Levels)
	case e.Message == "":
		return fmt.Errorf("%w: message is empty", ErrInvalidEvent)
	case e.At.IsZero():
		return fmt.Errorf("%w: at is the zero time", ErrInvalidEvent)
	}
	return nil
}
