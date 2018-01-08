package main

import (
	"fmt"
	"testing"
)

func BenchmarkAddSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	params := Params{}
	for i := 0; i < b.N; i++ {
		params.AddSlice(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
		params.AddSlice("foo", fmt.Sprintf("%d", i))
	}
}
func BenchmarkAddMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	params := MapParams{}
	for i := 0; i < b.N; i++ {
		params.AddMap(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
		params.AddMap("foo", fmt.Sprintf("%d", i))
	}
}
