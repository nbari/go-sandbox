package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := &Trie{}

	path := []string{"/"}
	trie.Set(path, nil, "ALL", "")

	path = []string{"/zero"}
	trie.Set(path, nil, "ALL", "")

	path = []string{"/", "foo"}
	trie.Set(path, nil, "ALL", "")

	path = []string{"/", "bar", ""}
	trie.Set(path, nil, "ALL", "")

	fmt.Printf("trie = %+v\n", trie)
	for _, n := range trie.Node {
		fmt.Printf("n = %+v\n", n)
		for _, nn := range n.Node {
			fmt.Printf("nn = %+v\n", nn)
		}
	}
}
