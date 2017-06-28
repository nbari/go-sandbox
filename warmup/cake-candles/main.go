package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var c, n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Fscan(io, &c)
		if val, ok := m[c]; !ok {
			m[c] = 1
		} else {
			m[c] = val + 1
		}
	}

	s := map[int]int{}
	keys := []int{}
	for k, v := range m {
		s[v] = k
		keys = append(keys, v)
	}
	sort.Ints(keys)
	// get last key keys[len(keys)-1] from map s and then from map m
	fmt.Printf("%d\n", m[s[keys[len(keys)-1]]])
}
