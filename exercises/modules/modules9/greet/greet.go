// Package greet builds greetings. The local copy adds Formal, which the
// published v1.0.0 does not have yet.
package greet

// Formal greets someone politely.
func Formal(name string) string {
	return "Good day, " + name + "."
}
