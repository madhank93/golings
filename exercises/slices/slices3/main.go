// slices3
// Make me compile!
//
// Add elements to a slice with append.

package main

import "fmt"

func main() {
	names := []string{"John", "Maria", "Carl", "Peter"}
	names = append(names, "Mary")
	fmt.Println(names)
}
