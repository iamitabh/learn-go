package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hfwjn48ied"

func main() {

	fmt.Println("Welcome to handling of URLs in Golang")
	fmt.Println(myurl)

	// In order to extract information from urls, we need to parse it first
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// A better way to extract query from the urls is as follow
	qparams := result.Query()
	// The type of output of Query() function is a slice of all the params
	fmt.Printf("The type of qparams is %T\n", qparams)

	fmt.Println(qparams["coursename"])

	// looping through that slice
	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	// How to create a url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
