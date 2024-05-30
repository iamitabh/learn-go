package main

import "fmt"

func main() {

	fmt.Println("Maps in golang")

	// Way to create a Map using "make" function
	language := make(map[string]string)
	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages: ", language)
	fmt.Println("JS shorts for: ", language["JS"])

	// In order to delete values from Map
	delete(language, "RB")
	fmt.Println("List of all languages: ", language)

	// How to loop in maps
	for key, value := range language {
		fmt.Printf("For key %v the value is %v\n", key, value)
	}

	// with comma ok syntax
	for _, value := range language {
		fmt.Printf("For key v the value is %v\n", value)
	}

}
