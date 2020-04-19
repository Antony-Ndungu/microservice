package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var writer = ioutil.Discard

func init() {
	go server()
}

func BenchmarkZero(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Post("http://localhost:8080/hello-world", "application/json", bytes.NewBuffer([]byte(`{"name": "Antony"}`)))
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Fatal(resp.StatusCode, err)
		}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&request)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkZeroUnMarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Post("http://localhost:8080/hello-world-unmarshal", "application/json", bytes.NewBuffer([]byte(`{"name": "Antony"}`)))
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Fatal(resp.StatusCode, err)
		}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&request)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkFirst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b, _ := json.Marshal(response)
		writer.Write(b)
	}

}

func BenchmarkSecond(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkThird(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}
