package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling of golang")
	content := "This need to go in a file"

	file, err := os.Create("./myLog.txt")

	if err != nil {
		// "panic" is a keyword used to stop the execution of file
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./myLog.txt")
}

func readFile(filename string) {
	// Whenever you read a file from disk, it is not in the format of string
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file in raw format is \n", databyte)
	// In order to change it to string, put that byte variable into string()
	fmt.Println("Text data inside the file in string format is \n", string(databyte))
}
