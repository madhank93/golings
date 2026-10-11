// modules8
// The email channel registers itself in its init function — but a package's
// init only runs if the package is part of the program. Make sure it is,
// without calling anything from it.

// I AM NOT DONE
package notify

import (
	"fmt"

	"example.com/notify/sender"
)

// Send delivers msg over the named channel.
func Send(channel, to, msg string) (string, error) {
	f, ok := sender.Lookup(channel)
	if !ok {
		return "", fmt.Errorf("no sender registered for %q", channel)
	}
	return f(to, msg), nil
}
