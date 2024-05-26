package main

import "fmt"

func main() {

	x := 0

	// While loop of go
	for x < 5 {
		fmt.Println("Value of x is :", x)
		x++
	}

	// For loop of go
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	// Looping through a slice
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
		// Remeber change done inside the bracket doesn't change the original slices
	}
	// If you only need to use value and not index and follow this syntax
	for _, value := range names {
		fmt.Println(value)
	}
}
