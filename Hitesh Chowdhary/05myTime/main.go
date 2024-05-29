package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// Notice very carefully that when formatting the time you alway alway need to
	// use "01-02-2006 15:04:05 Monday" only.
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// To create a Date
	createdDate := time.Date(2020, time.August, 10, 23, 23, 4, 6, 0, time.UTC)
	fmt.Println(createdDate)
}
