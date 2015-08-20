package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"strings"
)

func main() {
	trie := violetear.NewTrie()

	trie.Set([]string{""}, "root", "GET")
	trie.Set([]string{":uuid", ":uuid2"}, "h_uuid", "ALL")
	trie.Set([]string{":uuid"}, "h_uuid", "ALL")
	trie.Set([]string{"hello", "world", ":uuid", "cuatro"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hello", "world", ":uuidXY"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hello", "world", "last", ":uuid"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hello", "world", "last", "cuad"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hello", "world", "last", "cuatro"}, "h_hello_world", "GET, POST,     PUT")
	trie.Set([]string{"hello", "world", "last"}, "h_hello_world_sopas", "DELETE")
	trie.Set([]string{"hello", "world", "last"}, "h_hello_world_sopas_default", "OPTIONS")
	trie.Set([]string{"hello", "world"}, "h_hello_world", "GET")
	trie.Set([]string{"hello"}, "h_hello", "GET")
	trie.Set([]string{"hola", ":ip", ":uuid"}, "h_uuid", "ALL")
	trie.Set([]string{"hola", ":ip", "last"}, "h_uuid", "ALL")
	trie.Set([]string{"hola", ":uuid", ":last"}, "h_uuid", "ALL")
	trie.Set([]string{"hola", ":uuid", "last"}, "h_uuid", "ALL")

	r, l := trie.Get([]string{"hell3o"})

	if len(r.Handler) > 0 && l {
		fmt.Println(r.Handler, l)
	} else if r.HasRegex {
		// search for regex
		fmt.Println("----> ", r, r.HasRegex, l)
		for k, _ := range r.Node {
			if strings.HasPrefix(k, ":") {
				fmt.Println(k)
			}
		}

	} else {
		fmt.Println("no path", r)
	}

}
