package main

import (
	"fmt"
	"strings"
)

// How to return multiple value in a function
func getInitialis(n string) (string, string) {

	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {

	fn, sn := getInitialis("tifa lockhart")
	fmt.Println(fn, sn)

}
