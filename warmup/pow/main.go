package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	values := make([]int, 2)
	for i := range values {
		fmt.Fscan(io, &values[i])
	}
	out := 1
	for i := 0; i < values[1]; i++ {
		out = out * values[0]
	}
	fmt.Printf("out: %d\n", out)
}
