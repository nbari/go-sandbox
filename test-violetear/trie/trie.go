package main

import (
	"fmt"
	"github.com/nbari/violetear"
)

func main() {
	trie := violetear.NewTrie()
	fmt.Println(trie)

	m := map[string]string{"POST": "h_post", "GET": "h_get"}

	trie.Set([]string{""}, map[string]string{"ALL": "root"})
	trie.Set([]string{"hello"}, map[string]string{"POST": "h_hello"})
	trie.Set([]string{"hello", "world"}, map[string]string{"GET": "h_hello_world"})
	trie.Set([]string{"hello", "world", "last"}, m)
	trie.Set([]string{"hello", "world", "last"}, m)

	r, e := trie.Get([]string{"hello", "world", "last"})

	fmt.Println(r, e)

	//	log.Fatal(http.ListenAndServe(":8080", router))
}
