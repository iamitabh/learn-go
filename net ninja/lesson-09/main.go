package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good moring %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// Function can take another function as a argument
func cycleName(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}

}

// For returning value, you need to tell in the function definition itself
func circleArea(r float64) float64 {

	return math.Pi * r * r

}

func main() {

	sayGreeting("mario")
	sayBye("luigi")

	cycleName([]string{"Ckid", "Tifa", "barret"}, sayGreeting)

	fmt.Println(circleArea(10))
}
