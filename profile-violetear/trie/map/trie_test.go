package main

import (
	"testing"

	"github.com/nbari/violetear"
)

func BenchmarkTrie(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie := violetear.NewTrie()
		trie.Set([]string{"/"}, nil, "ALL", "v3")
	}
}
