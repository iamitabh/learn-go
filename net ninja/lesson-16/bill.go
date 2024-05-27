package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0.0,
	}

	return b
}

// receiver function: which is used to associate a function to a struct
func (b bill) format() string {
	fs := "Breakdown of bill \n"
	var total float64 = 0

	for k, v := range b.items {

		// Here note we have added -25 in "%-25v", that we have done so that this
		// part of string take 25 character on the right
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%v \n", "Total:", total)
	return fs
}
