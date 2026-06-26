// variables4
// Make me compile!
//
// See what happens when a block reuses a name with a new type.

// I AM NOT DONE
package main

import "fmt"

func main() {
	x := "TEN" // Don't change this line
	fmt.Printf("x has the value %s\n", x)

	if true {
		x = 1
		fmt.Println(x + 1)
	}

	fmt.Println(x)
}
