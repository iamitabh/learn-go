package main

import "fmt"

func main() {

	// Map is a data type in Golang similar to dicitonery in python. But the difference is that
	// datatype of key and datatype of value both should be consistent in all the data memeber in a map.

	menu := map[string]float64{
		"soup": 4.55,
		"pie":  84.3,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//You can change the values of maps
	phonebook := map[int]string{
		748393: "Luigi",
	}
	phonebook[748393] = "Mario"
}
