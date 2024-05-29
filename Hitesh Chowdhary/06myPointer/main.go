package main

import "fmt"

func main() {

	fmt.Println("Welcome to a class on pointer")

	// By default the value of pointer are "nil"
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	// Here we are refering the "myNumber" using &
	var ptr *int = &myNumber
	// Pointer are direct refer of memory location of that variable
	fmt.Println("Value of myNumber pointer: ", ptr)
	// Pointer with * is for checking the value in the pointer
	fmt.Println("Value of myNumber pointer: ", *ptr)

}
