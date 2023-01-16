package main

import (
	"fmt"
	"myapp/doctor"
)

func main() {
	// Prints, then inserts line break
	// fmt.Println("Hello, world!")

	// Prints both of these without a line break
	// fmt.Print("This is some text")
	// fmt.Print("this is some more text")

	// This is how to pass an argument to a func
	// sayHelloWorld(("Passing a string, as an argument to a func"))

	// Instead using a variable argument
	// var whatToSay string
	// whatToSay = "string from a var"
	// sayHelloWorld((whatToSay))

	// variable shorthand, assignment operator.
	// whatToSay := "create and store a var in one line, and determine its type automatically"
	// sayHelloWorld((whatToSay))

	// set a variable to the result of another method
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}

// Passing arguments to function.  Must list arg name and type.
// func sayHelloWorld(whatToSay string) {
// fmt.Println(whatToSay)
// }
