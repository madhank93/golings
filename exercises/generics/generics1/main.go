// generics1
// Make me compile!
//
// Write one function that works for any type.

package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

func print[T any](value T) {
	fmt.Println(value)
}
