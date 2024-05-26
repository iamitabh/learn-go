package main

import "fmt"

func updateName(x string) {
	x = "wedge"

	// in order to get the changes made in the function, you'd have to return the modified value
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {

	// Golang is a Pass by value language which means that Go makes a copies of values when passed into func.

	// This group value doesn't get updated inside function
	// group A type -> string, ints, bools, floats, arrays, structs
	name := "tifa"
	updateName(name)
	fmt.Println(name)

	// This group values does get updated
	// group B type -> slices, map, functions
	menu := map[string]float64{
		"pie":        5.67,
		"ice creame": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
