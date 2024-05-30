package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to our slices: ")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// How to add element to the slices
	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	// In the case you want to remove a set of elements from the slice
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// using make() function to define slice
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 942
	highScores[2] = 532
	highScores[3] = 832
	// highScores[4] = 777 //This will throw error. So, you can't add element like this way.
	// But when you use "append" function to add element then it's fine. Because, "append" relocate the memory
	highScores = append(highScores, 555, 666, 323)

	// To sort integers
	sort.Ints(highScores)
	fmt.Println(highScores)
	// To check if slice is sorted
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index
	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
