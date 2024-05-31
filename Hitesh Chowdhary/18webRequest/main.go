package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	// Notice, this line is very important to do.
	defer response.Body.Close()

	// ioutil.ReadAll(response.Body)
	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)

}
