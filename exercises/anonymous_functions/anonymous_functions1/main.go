// anonymous functions1
// Make me compile!

package main

import "fmt"

func main() {
	name := "Guorong"
	func(name string) {
		fmt.Printf("Hello %s", name)
	}(name)

}
