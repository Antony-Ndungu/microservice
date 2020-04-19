package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	message = "Hello"
	port    = 8080
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

var request helloWorldRequest
var response = helloWorldResponse{Message: message}

func main() {
	server()
}

func server() {
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/hello-world-unmarshal", helloWorldUnMarshalHandler)
	log.Printf("Server is starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldUnMarshalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusBadRequest))), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &request)

	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusBadRequest))), http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	response = helloWorldResponse{Message: fmt.Sprintf("%v %v", message, request.Name)}
	err = encoder.Encode(&response)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusInternalServerError))), http.StatusInternalServerError)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusBadRequest))), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	response = helloWorldResponse{Message: fmt.Sprintf("%v %v", message, request.Name)}
	err = encoder.Encode(&response)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusInternalServerError))), http.StatusInternalServerError)
	}
}
