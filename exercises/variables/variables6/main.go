// variables6
// Make me compile!
//
// Constants cannot be reassigned after declaration.

package main

import "fmt"

var x = 10

func main() {
	fmt.Println(x)

	x = x + 1
	fmt.Println(x)
}
