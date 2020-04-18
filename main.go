package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	message = "Hello world!"
	port    = 8080
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

var response = helloWorldResponse{Message: message}

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)
	log.Printf("Server is starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&response)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
	}
}
