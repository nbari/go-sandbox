package main

import (
	"fmt"
)

func noSplit(path string) {
	var key string
	if path != "" {
		fieldStart := -1
		for i := 0; i < len(path); i++ {
			if path[i] == '/' {
				if i >= 0 {
					if fieldStart >= 0 {
						key = path[1:i]
						path = path[i:]
						break
					}
				}
			} else if path[i] == '*' {
				key = "*"
				path = ""
				break
			} else if fieldStart == -1 {
				fieldStart = i
			}
		}
	} else {
		key = "/"
	}
	if key == "" && path != "" {
		key = path[1:]
		path = ""
	}
	if path == "/" {
		path = ""
	}
	fmt.Printf("key= %s path= %s\n", key, path)
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
		noSplit(p)
		println()
	}
}
