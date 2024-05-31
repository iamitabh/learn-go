package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

// Syntax to add Methods to struct
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// Because we are not passing pointer to the struct, this methods won't change the struct
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func main() {

	fmt.Println("Structs in Methods")

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16, 12}
	fmt.Println(hitesh)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Println(hitesh)

}
