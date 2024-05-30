package main

import "fmt"

// How to define struct
type User struct {
	// Here we are keeping the name of variable because these variables will be accessible to everyone
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	fmt.Println("Structs in Golang")
	// No inheritance in Golang, No Super or Parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	// using %+v placeholder give more data about struct
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)

}
