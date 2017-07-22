package main

import (
	"fmt"
	"strconv"
)

func main() {
	h := "2B55"
	i, err := strconv.ParseInt(h, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("i = %+v\n", i)
	r := rune(i)

	fmt.Printf("r = %c\n", r)
	fmt.Printf("2b55: %c\n", '\u2b55')
	fmt.Println("\u2b55")
}
