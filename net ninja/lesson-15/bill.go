package main

// struct definition
type bill struct {
	name string
	item map[string]float64
	tip  float64
}

// function to create new struct using the definition
func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}

	return b
}
