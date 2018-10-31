package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	x := GetOutboundIP()
	fmt.Printf("x = %+v\n", x)

}

// GetOutboundIP getoutboundip
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "1.1.1.1:53")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if ip, ok := conn.LocalAddr().(*net.UDPAddr); ok && !ip.IP.IsLoopback() {
		return ip.IP
	}
	//	return localAddr.IP
	return conn.LocalAddr().(*net.UDPAddr).IP
}
