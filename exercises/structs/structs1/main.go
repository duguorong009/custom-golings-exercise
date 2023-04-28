// structs1
// Make me compile!
package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func main() {
	person := Person{name: "John", age: 21}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
