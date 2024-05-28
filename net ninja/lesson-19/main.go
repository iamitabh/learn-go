package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Create a new bill name: ")
	// // This function captures the input from terminal till it's encounter \n or enter.
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func promptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("What's the item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)

	case "t":
		tip, _ := getInput("Enter the tip amount ($): ", reader)
		fmt.Println(tip)

	case "s":
		fmt.Println("You choose s")

	default:
		fmt.Println("You choose not a valid option...")
		promptOptions(b)
	}

}

func main() {

	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)

}
