package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 1
	u1 := uuid.NewV1()
	fmt.Printf("UUIDv1: %s\n", u1)

}
