package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i, input, n, positive, negative, zero float64
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	for i = 0; i < n; i++ {
		fmt.Fscan(io, &input)
		if input < 0 {
			negative++
		} else if input > 0 {
			positive++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f\n", positive/n, negative/n, zero/n)
}
