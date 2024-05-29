package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please rate our Pizza")
	fmt.Println("Please rate it between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating ", input)

	// strconv is a package used to convert strings
	// Note that we are writing "strings.TrimSpace(input)" because in "input" variable when we press enter
	// we get "\n" also in it. So, "\n" won't get converted to "float64". That's why we need to trim that part.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
