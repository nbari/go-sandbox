package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"strings"
)

func main() {
	trie := violetear.NewTrie()

	trie.Set([]string{"hello", ":uuid", "last"}, "h_uuid", "ALL")

	r, l := trie.Get([]string{"hello"})

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
