package main

import "fmt"
import "time"

func foo() {
	panic("no recovery")
}

func main() {
	go foo()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}()
	<-time.After(10 * time.Second)
}
