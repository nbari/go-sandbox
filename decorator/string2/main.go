package main

import (
	"fmt"
	"strings"
)

type Decorator func(string) string

func Upper() Decorator {
	return func(s string) string {
		return strings.ToUpper(s)
	}
}

func Prefix(prefix string) Decorator {
	return func(s string) string {
		return prefix + s
	}
}

func main() {
	decorators := []Decorator{
		Upper(),
		Prefix("MyPrefix"),
	}

	s := "-TheString"
	for _, dc := range decorators {
		s = dc(s)
	}

	fmt.Println(s)
}
