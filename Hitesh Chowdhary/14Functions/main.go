package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	// Writing just "greeter" means you are passing reference of the function
	greeter()

	result := adder(3, 2)
	fmt.Println("Result is :", result)

	// Note you can't define function inside function
}

// function signature of a function is the datatype you expect from the function to return
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// When you need to pass n number of parameters to a function
// values here will be accessible inside the function as a slice
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi the total is"
}

// // Anonymous functions
// func () {
// 	fmt.Println("Anonymous functions")
// } ()

func greeter() {
	fmt.Println("Namstey from Golang")
}
