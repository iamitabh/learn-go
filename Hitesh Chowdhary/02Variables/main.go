package main

import "fmt"

// Here, you can't (In global scope) declare variable using walrus operator
// jwtToken := 304i2fnf0ijd

// Declaration of const
const LoginToken string = "hjanjanfkfs"

// Notic, "L" of "LoginToken" is capital that means in GoLang that "LoginToken" is publicly avaiable

func main() {

	var username string = "amitabh"
	fmt.Println(username)
	// Printf is used to print out formatted string
	// Here, we are putting "%T" placeholder, this "%T" placeholder tells the type of variable
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// Similar to "uint8" we also have "uint16", "uint32", "uint64"
	// Because we have so many option, instead of these you could direct choose "uint"
	// Here, "u" stands for un-signed integer
	var smallVal uint8 = 255 //Can't put interger values greater then 255
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// Similar story to "uint" for "int" also.

	// We have "float64" also.
	// You can choose just "float" datatype also.
	var smallFloat float32 = 255.43545 //In float32, after decimal you get only values till 5 places.
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default value of an variable
	var anotherVariable int
	fmt.Println(anotherVariable) //Here we will get 0
	// Golang doesn't put garbage values in the variables

	//Implicit way of declarating variable (Where you don't put type of variable while declaring)
	var website = "google.com"
	fmt.Println(website)
	// Here, the golang will automatically get the type (in this example string) and you won't be
	// able put "int", "float" or any other variable values in the "website" variable.

	// short variable declaration style by using walrus operator
	numberOfUser := 300000
	fmt.Println(numberOfUser)

}
