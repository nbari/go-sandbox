package main

import (
	"bufio"
	"fmt"
	"os"
)

// cat test.txt| go run main.go
func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(io, &v)
		a[i] = v
	}

	for k, v := range a {
		fmt.Printf("k=%d v=%d\n", k, v)
	}
}
