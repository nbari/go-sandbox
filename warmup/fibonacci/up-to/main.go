package main

import "fmt"

const finish = 100

func fibonacci(m, n int) {
	if n >= finish {
		return
	}
	fmt.Println(n)
	fibonacci(n, (m + n))
}

func main() {
	fibonacci(0, 1)
}
