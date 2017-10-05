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
		if path[i] == '/' {
			if i == 0 {
				return noSplit(path[1:])
			}
			if i > 0 {
				key = path[:i]
				path = path[i:]
				if key == "" && path != "" {
					return noSplit(path)
				}
				if path == "/" {
					return key, ""
				}
				return key, path
			}
		}
	}
	return path, ""
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
		"/hello/",
	}
	for _, p := range paths {
		fmt.Printf("p = %+v\n", p)
		key, path := noSplit(p)
		fmt.Printf("key: %s path: %s\n", key, path)
		println()
	}
}
