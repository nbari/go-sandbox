package main

import "testing"

func BenchmarkA(b *testing.B) {
	A("switch vs if")
}

func BenchmarkAalpha(b *testing.B) {
	A("alpha")
}

func BenchmarkB(b *testing.B) {
	B("switch vs if")
}

func BenchmarkBalpha(b *testing.B) {
	B("alpha")
}
