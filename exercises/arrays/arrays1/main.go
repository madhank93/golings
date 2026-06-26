// arrays1
// Make me compile!
//
// Arrays are zero-indexed — fix the out-of-range access.

// I AM NOT DONE
package main

import "fmt"

func main() {
	var colors [3]string

	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Printf("First color is %s\n", colors[])
	fmt.Printf("Last color is %s\n", colors[])
}
