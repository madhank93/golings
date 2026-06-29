// slices2
// Make me compile!
//
// Take a sub-slice using [low:high] bounds.

package main

import "fmt"

func main() {
	names := [4]string{"John", "Maria", "Carl", "Peter"}
	lastTwoNames := names[2:]
	fmt.Println(lastTwoNames)
}
