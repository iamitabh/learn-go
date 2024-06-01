package main

// go build .
// go mod tidy
// go mod verify
// go list
// go list -m all
// go list -m -versions github.com/gorilla/mux
// go why github.com/gorilla/mux
// go mod graph
// go mod edit -go 1.16
// go mod edit -module 10
// go mod vendor && go run -mod=vendor main.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}
