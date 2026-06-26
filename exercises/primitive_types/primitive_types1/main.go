// primitive_types1
// Make me compile!
//
// Toggle a boolean before the second check.

// I AM NOT DONE
package main

import "fmt"

func main() {
	storeIsOpen := true
	if storeIsOpen {
		fmt.Println("The store is open, let's buy some clothes!")
	}

	storeIsOpen
	if !storeIsOpen {
		fmt.Println("Oh no, let's buy some clothes online!")
	}
}
