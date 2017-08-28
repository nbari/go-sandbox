package main

import "fmt"

// unique pairs (5,0), (3,2), (1,4)
func main() {
	nums := []int{5, 3, 7, 0, 1, 4, 2}
	s := 5

	m := map[int]struct{}{}

	for _, v := range nums {
		if _, ok := m[s-v]; ok {
			fmt.Printf("pairs(%d, %d)\n", s-v, v)
		} else {
			m[v] = struct{}{}
		}
	}
}
