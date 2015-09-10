package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"net/http"
	"strings"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "aa")
}

func main() {
	trie := violetear.NewTrie()

	trie.Set([]string{"hello", ":uuid", "last"}, test, "ALL")

	t, p, l := trie.Get([]string{"hello"})

}
