package main

import "fmt"

func main() {

	//strings
	//var nameOne string = 45 //Error
	var nameOne string = "mario" //If you have declared a variable and not used still it'd give error
	var nameTwo = "lugi"         //If you don't specify the type of variable explicitly, it'll infer automatically
	var nameThree string         //Only we're declarating here, so here we'll get empty string in this variable

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	//Another syntax to initializing/declarating variables

	nameFour := "Yoshi" //You can't use this syntax outside of the function block
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// We can specify the bit size of int also
	// bits & memory
	var numOne int8 = 25
	// we have int8, int16, int32, int64
	var numTwo int8 = -128

	//unsign int
	var numThree uint16 = 25 //You can't put negative number in these

	// floats -> in floats you need to specify the number of bits needed unlike int
	var scoreOne float32 = 25.84
	var scoreTwo float64 = 4820853802932048532.5892420
	var scoreThree := 1.5 //This will create float64
}
