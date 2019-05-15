package main

import (
	"fmt"
)

var seq = []int{0, 1}

func fib(n int) int {
	if n < 2 {
		return n
	}
	seq = append(seq, fib(seq[n-1])+fib(seq[n-2]))
	fmt.Printf("seq = %+v\n", seq)
	fmt.Printf("n = %+v\n", n)
	if n >= 10 {
		return n
	}
	return fib(n + 1)
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	fmt.Println(fib(2))
}
