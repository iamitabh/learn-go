package main

import "fmt"

func main() {

	// Initialization of Array. Array in Go have fixed lenght.
	//var ageOne [3]int = [3]int{25, 30, 42}
	var ageTwo = [3]int{23, 49, 32}
	name := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	name[1] = "lugi" //You can access array element and manipulate them

	fmt.Println(ageTwo, len(ageTwo))
	fmt.Println(name, len(name))

	// Slices mean array with flexiable length
	var score = []int{100, 50, 60}
	score[2] = 34
	score = append(score, 45) //Only work on SLices not on array and return new slice and doesn't change the given slice

	// Slice ranges
	rangeOne := name[1:3]
	fmt.Println(rangeOne)
}
