// functions3
// Make me compile!
//
// Call a function with the argument it expects.

package main

import "fmt"

func main() {
	call_me(10)
}

func call_me(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
