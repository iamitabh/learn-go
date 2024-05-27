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
		items: map[string]float64{},
		tip:   0.0,
	}

	return b
}

// receiver function: which is used to associate a function to a struct
func (b *bill) format() string {
	fs := "Breakdown of bill \n"
	var total float64 = 0

	for k, v := range b.items {

		// Here note we have added -25 in "%-25v", that we have done so that this
		// part of string take 25 character on the right
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "Tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%v \n", "Total:", total+b.tip)
	return fs
}

// Update the tip data memeber
func (b *bill) UpdateTip(tip float64) {
	// here, we are de-referencing the pointer the by using (*b)
	// But, in struct of golang you actually don't need to de-referenc the pointer
	(*b).tip = tip
}

// Add an item to the bill
func (b *bill) AddItem(name string, price float64) {
	// Like here we are not de-referecing but still it'd work. You just need to associate the pointer of
	// struct to the function correctly
	b.items[name] = price
}
