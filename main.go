package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/hello-world", newValidationHandler(newHelloWorldHandler()))
	http.HandleFunc("/hello-world-unmarshal", helloWorldUnMarshalHandler)
	log.Printf("Server is starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) *validationHandler {
	return &validationHandler{next: next}
}

func (v *validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s", r.URL)
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusBadRequest))), http.StatusBadRequest)
		return
	}
	v.next.ServeHTTP(w, r)
}

type helloWorldHandler struct {
}

func newHelloWorldHandler() *helloWorldHandler {
	return &helloWorldHandler{}
}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	response = helloWorldResponse{Message: fmt.Sprintf("%v %v", message, request.Name)}
	err := encoder.Encode(&response)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": %v }", fmt.Sprintf("\"%v\"", http.StatusText(http.StatusInternalServerError))), http.StatusInternalServerError)
	}
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

func fetchGoogle() {
	r, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancel()

	r = r.WithContext(ctx)

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Status)
}
