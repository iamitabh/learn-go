package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {

	name := "tifa"

	// You can access the memory address by using &
	fmt.Println("Memory address of name variable is ", &name)

	//Here, m is variable which store other variables memory address and is called pointer
	m := &name
	fmt.Println("pointer m ", m)

	//We can access the value store in that memory address by using *
	fmt.Println("value stored at the memory address", *m)

	// Now because we passing the memory address of variable name, so we can change it value inside function
	updateName(&name)
	fmt.Println(name)
}
