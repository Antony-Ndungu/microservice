package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var writer = ioutil.Discard

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
