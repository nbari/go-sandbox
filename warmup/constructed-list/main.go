package main

import "fmt"

func main() {
	println(Solution([]int{1, 4, -1, 3, 2}))
}

func Solution(A []int) int {
	var l, k, i int
	for j := 0; j < len(A); j++ {
		k = A[i]
		fmt.Printf("k = %+v\n", k)
		if k < 0 {
			l++
			return l
		}
		l++
		i = k
	}
	return l
}
