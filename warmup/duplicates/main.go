package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 1, 5, 9, 3, 5, 6, 0}
	m := make(map[int]int)
	out := []int{}
	for _, v := range x {
		if _, ok := m[v]; ok {
			continue
		} else {
			m[v] = 0
			out = append(out, v)
		}
	}
	fmt.Printf("%v\n", out)
}
