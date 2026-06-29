// switch2
// Make me compile!
//
// A conditionless switch needs a boolean case or default.

package main

import "fmt"

func main() {
	switch {
	case 0 > 1:
		fmt.Println("zero is greater than one")
	case 1 > 0:
		fmt.Println("one is greater than zero")
	}
}
