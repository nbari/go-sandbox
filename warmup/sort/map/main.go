package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "abcccccccd"
	m := make(map[string]int)
	for _, s := range str {
		if val, ok := m[string(s)]; ok {
			m[string(s)] = val + 1
		} else {
			m[string(s)] = 1
		}
	}
	n := map[int]string{}
	var a []int
	for k, v := range m {
		n[v] = k
		a = append(a, v)
	}

	sort.Ints(a)

	fmt.Printf("n[a[len(a)-1]] = %+v\n", n[a[len(a)-1]])
}
