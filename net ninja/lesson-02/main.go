package main

// In Go, main is actually a special package name which indicates that the package contains the code for
// an executable application. That is, it indicates that the package contains code that can be built
// into a binary and run.

// Like if you'd be making some utils program or shared library then you wouldn't put "package main"

import "fmt"

// Here we are importing package for formating strings and printing messages

// All function start with "func" keyword and the name of this function is important "main" because
// it is a starting point of our program. There should be only one "main" function in the entire app.
func main() {

	// The methods we use from the imported files starts with Captials like "P" of "Println".
	fmt.Println("Hello, ninjas!")
}
