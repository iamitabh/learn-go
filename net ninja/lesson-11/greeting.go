package main

import "fmt"

var points = []int{20, 24, 58, 194}

// Any function you declare here, will atomatically be accessible to the main.go file because both files
// are part of the same "package main"
func sayHello(n string) {
	fmt.Println("hello", n)
}
