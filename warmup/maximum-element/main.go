package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	var n, q int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	s := make(stack, 0)
	for i := 0; i <= n; i++ {
		fmt.Fscan(io, &q)
		switch q {
		case 1:
			fmt.Fscan(io, &q)
			s = s.Push(q)
		case 2:
			s, _ = s.Pop()
		case 3:
			s, q = s.Pop()
			fmt.Println(q)
		}
	}
}
