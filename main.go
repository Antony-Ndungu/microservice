package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	message = "Hello world!"
	port    = 8080
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Printf("Server is starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v\n", message)))
}
