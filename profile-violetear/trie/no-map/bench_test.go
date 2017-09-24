package trie

import (
	"fmt"
	"testing"
)

func BenchmarkTrie(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	trie := &Trie{}
	path := []string{"/"}

	for i := 0; i < b.N; i++ {
		trie.Set(path, nil, "ALL", "v3")
		trie.Set(path, nil, "ALL", "")
		trie.Get(path, "v3")
		path = append(path, fmt.Sprintf("%d", i))
	}
}
