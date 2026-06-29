// range2
// Make me compile!
//
// Iterate a map's keys and values with range.

package main

import "fmt"

func main() {
	phoneBook := map[string]string{
		"Ana":  "+01 101 102",
		"John": "+01 333 666",
	}

	for name, phone := range phoneBook {
		fmt.Printf("%s has the %s phone\n", name, phone)
	}
}
