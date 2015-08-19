package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"strings"
)

func main() {
	trie := violetear.NewTrie()

	trie.Set([]string{""}, "root", "GET")
	trie.Set([]string{"hello"}, "h_hello", "GET")
	trie.Set([]string{"hello", "world"}, "h_hello_world", "GET")
	trie.Set([]string{"hello", "world", "last"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hello", "world", "last"}, "h_hello_world_sopas", "DELETE")
	trie.Set([]string{"hello", "world", "last"}, "h_hello_world_sopas_default", "OPTIONS")
	trie.Set([]string{"hola", "world", "last", "cuatro"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hola", ":uuid", "last"}, "h_uuid", "ALL")

	//	l, r := trie.Get([]string{"hola", ":uuid", "last"})
	//l, r := trie.Get([]string{"xhello"})
	r := trie.Get([]string{"hola", "sopas"})
	//r := trie.Get([]string{"hola", "world", "last", "cuatro"})

	//	l, r := trie.Get([]string{"hello", "world", "last"})

	if len(r.Handler) > 0 {
		fmt.Println(r.Level, r.Handler)
	} else {

		//	trie.GetLevel(3)

		for k, _ := range r.Node {
			if strings.HasPrefix(k, ":") {
				fmt.Println(k)
			}
		}

	}

}
