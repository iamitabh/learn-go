package main

import "fmt"

func main() {
	age := 34
	name := "Shaun"

	// Print
	fmt.Print("hello, ") //Doesn't add new line

	// Print with new line
	fmt.Println("Hello, Ninja")
	fmt.Println("My name is", name, "and my age is", age)

	// Printf (formatted strings) and this doesn't add new line
	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v is a format specifier
	fmt.Printf("my age is %q and my name \n", name)            // %q is a format specifier that puts quotes on strings
	fmt.Printf("age is of type %T \n", age)                    // %T tells the type of the variable
	fmt.Printf("you scored %0.1f points \n", 225.332)          // %f is for float values

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is : ", str)
}
