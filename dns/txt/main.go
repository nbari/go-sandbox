package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	txt, err := net.LookupTXT("o-o.myaddr.l.google.com")
	if err != nil {
		txt, err := net.LookupTXT("whoami.akamai.net")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(txt)
	}
	fmt.Println(txt)
}
