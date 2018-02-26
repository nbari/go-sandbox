package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	domain := flag.String("domain", "stackoverflow.com", "domain name")
	flag.Parse()
	fmt.Printf("domain %s\n", *domain)

	server := "8.8.4.4"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(*domain+".", dns.TypeA)
	r, _, err := c.Exchange(&m, server+":53")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(r.Answer) == 0 {
		fmt.Println("Could not found NS records")
		os.Exit(1)
	}

	for _, ans := range r.Answer {
		if a, ok := ans.(*dns.A); ok {
			fmt.Printf("%s. %d IN A %s\n", *domain, ans.Header().Ttl, a.A.String())
		}
	}
}
