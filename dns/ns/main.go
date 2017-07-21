package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	domain := flag.String("domain", "immortal.run", "domain name")
	flag.Parse()
	fmt.Printf("domain %s\n", *domain)
	ns, err := net.LookupNS(*domain)
	if err != nil {
		fmt.Printf("err = %s\n", err)
		os.Exit(1)
	}
	for _, k := range ns {
		fmt.Printf("ns = %s\n", k.Host)
	}
}
