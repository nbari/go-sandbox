package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	fmt.Fscan(io, &s)
	var sea, valleys int
	for i := 0; i < n; i++ {
		if string(s[i]) == "U" {
			sea++
		} else if string(s[i]) == "D" {
			sea--
		}
		if sea == 0 && string(s[i]) == "U" {
			valleys++
		}
	}
	fmt.Println(valleys)
}
