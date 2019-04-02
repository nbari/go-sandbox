package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, s int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	set := make(map[int]int)
	pairs := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &s)
		if _, ok := set[s]; !ok {
			set[s] = 1
		} else {
			pairs++
			delete(set, s)
		}
	}
	fmt.Println(pairs)
}
