package main

// Greet lives in a second file on purpose: a package-mode exercise must be
// compiled as a whole directory, so a runner that passes only main_test.go to
// `go test` fails here with "undefined: Greet".
func Greet(name string) string {
	return "hello, " + name
}
