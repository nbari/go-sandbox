package main

import (
	"fmt"
)

func fib(n int) int {
	fib := []int{0, 1}
	for i := 2; i <= n; i++ {
		n := fib[i-1] + fib[i-2]
		fib = append(fib, n)
	}
	return fib[n]
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	fmt.Println(fib(4))
}
