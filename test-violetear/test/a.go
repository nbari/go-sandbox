package main

import "fmt"

func Add(resource string, handler string, methods string) int {
	fmt.Println(resource, handler, methods)
	return 1
}

func main() {

	Add("resource", "handler", "GET, POST")

}
