package main

import (
	"bytes"
	"fmt"
)

func noSplit(path string) {
	var field bytes.Buffer
	for _, rune := range path {
		if rune == '/' {
			if field.Len() > 0 {
				fmt.Printf("field = %+v\n", field.String())
				field.Reset()
			}
			continue
		} else if rune == '*' {
			fmt.Printf("field = %c\n", rune)
			field.Reset()
			return
		}
		field.WriteRune(rune)
	}
	if field.Len() > 0 {
		fmt.Printf("field = %+v\n", field.String())
	}
}

func main() {
	paths := []string{
		"/hello",
		"/hello/world",
		"/hello/:uuid",
		"/hello/:uuid/:uuid",
		"/hello/:uuid/:uuid/*",
		"/hello/:uuid/:uuid/*/",
		"/hello/*/:uuid/:uuid/*/",
	}
	for _, p := range paths {
		fmt.Printf("p = %+v\n", p)
		noSplit(p)
	}
}
