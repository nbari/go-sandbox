package main

import "fmt"

type Foo struct {
	a          string
	b          string
	c          string
	serverName string
}

// mafin
func main() {
	foo := Foo{a: "11", b: "aa", c: "dd"}
	//f := func() {
	//fmt.Printf(" = %+v\n")
	//}

	fmt.Printf("foo = %+v\n", foo)
}
