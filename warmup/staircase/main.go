package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= n-i {
				print(" ")
			} else {
				print("#")
			}
		}
		println()
	}
}
