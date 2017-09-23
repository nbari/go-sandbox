package main

import (
	"github.com/nbari/violetear"
)

func main() {
	trie := violetear.NewTrie()
	trie.Set([]string{"/"}, nil, "ALL", "v3")
}
