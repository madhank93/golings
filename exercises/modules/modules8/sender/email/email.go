// Package email registers the "email" delivery channel. It exports nothing:
// importing it is how a program opts in.
package email

import "example.com/notify/sender"

func init() {
	sender.Register("email", func(to, msg string) string {
		return "mailed " + to + ": " + msg
	})
}
