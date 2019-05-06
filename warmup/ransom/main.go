package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, y int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &x)
	fmt.Fscan(io, &y)
	magazine := make(map[string]int, x)
	ransom := make([]string, y)

	for i := 0; i < x; i++ {
		var s string
		fmt.Fscan(io, &s)
		if val, ok := magazine[s]; ok {
			magazine[s] = val + 1
		} else {
			magazine[s] = 1
		}
	}

	for i := range ransom {
		fmt.Fscan(io, &ransom[i])
	}

	for _, v := range ransom {
		if val, ok := magazine[v]; ok {
			if val > 1 {
				val--
				magazine[v] = val
			} else {
				delete(magazine, v)
			}
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
