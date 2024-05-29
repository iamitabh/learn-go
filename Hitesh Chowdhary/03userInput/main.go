package main

import (
	"bufio"
	"fmt"
	"os"
)

// If you want to search of any specific package of Golang you can do that on pkg.go.dev website

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our Pizza: ")

	// comma error syntax or comma ok syntax

	// input, err - in case you want to deal with "err" part
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
	// Here, you will notice that the type of "input" is "string". So, we'd only able to do the "string"
	// related operation on it. Hence, we need to convert it to "int" in next lesson.
}
