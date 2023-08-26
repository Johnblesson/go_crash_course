package main

import "fmt"

var name = "Kharis"
func main() {
	// Using Var
	// var name = "Kharis"
	var age int32 = 37
	var isCool = true

	// Shorthand
	name := "John Blesson Rowe" // This method can only be inside of a function to run
	size := 1.3
	name, email := "John", "john@rowe.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}