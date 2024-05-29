package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// fruitList[2] is missing
	fruitList[3] = "Peach"

	// Here, at fruitList[2] you'd notice blank space
	fmt.Println("Fruits list is: ", fruitList)
	fmt.Println("Fruits list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushrooms"}
	fmt.Println("Veggy List is: ", vegList)
	fmt.Println("Veggy List is: ", len(vegList))

}
