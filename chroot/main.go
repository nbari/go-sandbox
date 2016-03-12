package main

import (
	"fmt"
	"syscall"
)

func main() {

	err := syscall.Chroot("/tmp/root")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

}
