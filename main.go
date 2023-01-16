package main

import "fmt"

func main() {
	// Prints, then inserts line break
	// fmt.Println("Hello, world!")
	// Prints both of these without a line break
	// fmt.Print("This is some text")
	// fmt.Print("this is some more text")
	// This is how to pass an argument to a func
	sayHelloWorld(("Passing a string, as an argument to a func"))
}

// Passing arguments to function.  Must list arg name and type.
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
