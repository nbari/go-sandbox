package main

import (
	"fmt"
)

func fib(n int) int {
	seq := []int{0, 1}
	for i := 2; i <= n; i++ {
		n := seq[i-1] + seq[i-2]
		seq = append(seq, n)
	}
	return seq[n]
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	fmt.Println(fib(4))
}
