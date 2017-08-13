package main

import "fmt"

// unique paris (5,0), (3,2), (1,4)
func main() {
	arr := []int{5, 3, 7, 0, 1, 4, 2}
	s := 5

	input := map[int]int{}

	for k, v := range arr {
		input[v] = k
	}

	for _, v := range arr {
		_, ok := input[s-v]
		if ok {
			fmt.Printf("pairs(%d, %d)\n", v, s-v)
		}
	}
}
