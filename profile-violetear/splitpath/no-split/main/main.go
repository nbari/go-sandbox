package main

import (
	"bytes"
	"fmt"
)

func noSplit(path string) {
	var field bytes.Buffer
	for i, rune := range path {
		if rune == '/' && i > 0 {
			fmt.Printf("field = %+v\n", field.String())
			fmt.Printf("path = %+v\n", path[i:])
			return
		} else if rune == '*' {
			fmt.Printf("field = %c\n", rune)
			return
		} else if rune != '/' {
			field.WriteRune(rune)
		}
	}
	if field.Len() > 0 {
		fmt.Printf("field = %+v\n", field.String())
		fmt.Printf("path = %+v\n", "")
	}
}

func main() {
	paths := []string{
		"/hello/:uuid/:uuid",
		"/:uuid/:uuid",
		"/:uuid",
	}
	for _, p := range paths {
		fmt.Printf("p = %+v\n", p)
		noSplit(p)
		println()
	}
}
