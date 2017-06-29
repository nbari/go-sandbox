package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	io := bufio.NewReader(os.Stdin)
	fmt.Fscanln(io, &s)
	ampm := s[len(s)-2:]
	t := strings.Split(s[:len(s)-2], ":")
	var t2 = []int{}
	for _, i := range t {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	if ampm == "AM" {
		if t2[0] == 12 {
			t2[0] = 00
		}
		fmt.Printf("%02d:%02d:%02d\n", t2[0], t2[1], t2[2])
	} else {
		var h int
		if t2[0] != 12 {
			h = t2[0] + 12
			if h == 24 {
				h = 23
			}
		} else {
			h = 12
		}
		fmt.Printf("%02d:%02d:%02d\n", h, t2[1], t2[2])
	}
}
