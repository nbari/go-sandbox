package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	domain := flag.String("domain", "immortal.run", "domain name")
	flag.Parse()
	fmt.Printf("domain %s\n", *domain)

	server := "84.200.70.40"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(*domain+".", dns.TypeNS)
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
		ns, ok := ans.(*dns.NS)
		if ok {
			fmt.Printf("ns: %s\n", ns.Ns)
		}
	}
}
