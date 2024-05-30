package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case in GoLang")

	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
		// fallthrough
		// "fallthrough" is used if you want to run every switch case after a particular case
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("Don't know")
	}
}
