package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greet := "Hello, there friends!"

	// These are the function avaiable in "string" package
	// These function don't alter the original string, they return new string
	fmt.Println(strings.Contains(greet, "Hello"))         // strings.Contains function check if the given substring exists or not
	fmt.Println(strings.ReplaceAll(greet, "Hello", "Hi")) //Replace given substring
	fmt.Println(strings.ToUpper(greet))                   //make upper case the whole string
	fmt.Println(strings.Index(greet, "ll"))               //finds the index of given substring
	fmt.Println(strings.Split(greet, " "))                //splits the strings in several parts in arrays

	ages := []int{45, 20, 35, 30, 64, 43, 59}

	// These are function avaiable in "sort" package
	sort.Ints(ages) //This function alters the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 40) //Returns the index of provided int
	fmt.Println(index)

	names := []string{"Yoshi", "Mario", "Peach", "Bowser", "Lugi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Lugi"))
}
