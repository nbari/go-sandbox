package main

import (
	"fmt"
	"testing"

	"github.com/nbari/violetear"
)

func BenchmarkTrie(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	trie := violetear.NewTrie()
	path := []string{"/"}

	for i := 0; i < b.N; i++ {
		trie.Set(path, nil, "ALL", "v3")
		path = append(path, fmt.Sprintf("%d", i))
	}
}
