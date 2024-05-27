package main

import "fmt"

func main() {

	mybill := newBill("Mario's bill")

	mybill.AddItem("Onion soup", 4.50)
	mybill.AddItem("Veg Pie", 8.95)
	mybill.AddItem("Toffee Pudding", 4.95)
	mybill.AddItem("Coffee", 3.25)
	mybill.UpdateTip(10)
	fmt.Println(mybill.format())

}
