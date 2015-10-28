package main

import "testing"

var rootsValid = []string{"foo", "bar", "qwe", "asd"}

var void struct{}
var rootsValid2 = map[string]struct{}{"foo": void, "bar": void, "qwe": void, "asd": void}

// uses a slice
func BenchmarkSearch1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range rootsValid {
			if v == "/ja" {
				break
			}
		}
	}
}

// uses a map
func BenchmarkSearch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := rootsValid2["/ja"]; ok {
			continue
		}
	}
}
