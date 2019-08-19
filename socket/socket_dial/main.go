package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
)

func main() {
	l, err := net.Dial("unix", filepath.Join("/tmp", "my.sock"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("l = %+v\n", l)
}
