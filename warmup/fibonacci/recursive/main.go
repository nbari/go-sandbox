package main

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	fmt.Println(fib(7))
}
