package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer")
	// "defer" keyword delays the execution of line at the last end and it's work in LIFO fashion.
	defer fmt.Println("One")
	defer fmt.Println("Two")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
