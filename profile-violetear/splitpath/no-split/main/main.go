package main

import (
	"fmt"
)

func noSplit(path string) (string, string) {
	var key string
	if path == "" {
		return key, ""
	}
	for i := 0; i < len(path); i++ {
		if path[i] == '/' && i > 0 {
			return path[1:i], path[i:]
		}
	}
	return path[1:], ""
}

func main() {
	paths := []string{
		"/",
		"//",
		"//////////",
		"/hello/:uuid/:uuid",
		"/hello/:uuid/:uuid/",
		"/:uuid/:uuid/",
		"/:uuid/",
		"/:uuid/:uuid",
		"/:uuid",
	}
	for _, p := range paths {
		fmt.Printf("p = %+v\n", p)
		key, path := noSplit(p)
		fmt.Printf("key: %s path: %s\n", key, path)
		println()
	}
}
