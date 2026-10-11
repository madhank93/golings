// Package sender is a registry of delivery channels. Channel packages add
// themselves from their init functions, the way database/sql drivers do.
package sender

// Func delivers msg to a recipient and describes what it did.
type Func func(to, msg string) string

var registry = map[string]Func{}

// Register makes a channel available under name.
func Register(name string, f Func) { registry[name] = f }

// Lookup returns the channel registered under name.
func Lookup(name string) (Func, bool) {
	f, ok := registry[name]
	return f, ok
}
