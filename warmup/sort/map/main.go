package main

import (
	"fmt"
)

func main() {
	str := "abcccccccd"
	//str := "apple 1231111"
	m := make(map[string]int)
	for _, s := range str {
		if val, ok := m[string(s)]; ok {
			m[string(s)] = val + 1
		} else {
			m[string(s)] = 1
		}
	}
	var c string
	var l int
	for k, v := range m {
		if v > l {
			c = k
			l = v
		}
	}

	fmt.Println(c)
}
