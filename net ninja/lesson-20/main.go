package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		//This is a function used to parse a string into a float value
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number:")
			promptOptions(b)
		}
		b.AddItem(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter the tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number:")
			promptOptions(b)
		}
		b.UpdateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)

	case "s":
		fmt.Println("You choose to save the bill\n", b)

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
