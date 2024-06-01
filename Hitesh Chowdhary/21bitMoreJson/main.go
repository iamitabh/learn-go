package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // The 3rd parameter in defining struct is kind of a alias which get used when we create a JSON using struct
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tag,omitempty"` //Don't put space after comma and before "omitempty" which removes the field if the given field is empty
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"React Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "LearnCodeOnline.in", "bcd123", []string{"full-dev", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// Package the above data in JSON
	// finalJson, err := json.Marshal(lcoCourses)
	// To make JSON more readable, use MarshalIndent func
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	// When you receive the data, it comes in format of slice/array of byte
	jsonDataFromWeb := []byte(`
	{
		"coursename":"ReactJS",
		"Price":299,
		"website":"LearnCodeOnline.in",
		"tags":["web-dev","js"]
	}
	`)

	var lcoCourses course

	// To check if JSON data is correct or not
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		// Remeber here pass the address of struct/interface because you'd want to save data in it not in it's copy somewhere else
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		// %#v is a placeholder for struct/interface
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON was not valid")
	}

	// In some case where you don't want to create a struct

	// the key will be string but we can't say about the value what datatype that'll be off
	var OnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &OnlineData)
	fmt.Printf("%#v\n", OnlineData)

	// Althought for the value part of key-value we have defined it as a interface, but when converting it to JSON
	// will automatically assign it a propert datatype which would make sense to it.
}
