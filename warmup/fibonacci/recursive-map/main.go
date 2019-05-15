package main

import (
	"fmt"
)

var cache = make(map[int]int)

func fib(n int) int {
	if val, ok := cache[n]; ok {
		println(val)
		return val
	} else {
		nn := slowFib(n)
		cache[n] = nn
		return nn
	}
}

func slowFib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	fmt.Println(fib(70))
}
