package main

import "fmt"

func main() {

	fmt.Println("If else in golang")
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Don't know"
	}

	fmt.Println(result)

	// A syntax which get used in handling web request
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// if err != nil {

	// }

}
