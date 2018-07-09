package basic

import "fmt"

const pi = 3.14

/*
Constants - Example of using const in Go

Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.
*/
func Constants() {
	fmt.Printf("\n****Running basic.Constants(), constants example ")
	const World = "世界"
	fmt.Println("\n Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
