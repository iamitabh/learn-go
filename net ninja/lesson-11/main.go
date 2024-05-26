package main

import "fmt"

// To make file accessible to other file of the package main, defined func, var here outside of the main func

func main() {

	// Direclty use this function
	sayHello("mario")

	// points is also not defined here, but it defined in greeting file
	for _, v := range points {

		fmt.Println(v)

	}

	// Note, if you don't run the both files (main.go, greeting.go) then you'd get a sayHello not found error.

}
