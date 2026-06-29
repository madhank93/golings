// anonymous functions1
// Make me compile!
//
// Pass an argument to an anonymous function.

package main

import "fmt"

func main() {

	func(name string) {
		fmt.Printf("Hello %s", name)
	}("Gopher")

}
