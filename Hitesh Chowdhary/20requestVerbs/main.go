package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request from Golang")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	// One way of handling repsonse and converting it to string
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println("Content of Request Body: ", string(content))

	// Second way of handling response and converting it to string
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("The response is: ", responseString.String())

}

func PerformPostRequest() {

	const url string = "http://localhost:3000/post"

	// How to create JSON in Golang
	requestBody := strings.NewReader(`
		{
			"courseName": "Golang programming"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))
}

func PerformPostFormRequest() {

	const myurl string = "http://localhost:3000/post"

	// How to create form-data
	data := url.Values{}
	data.Add("firstName", "Amitabh")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
