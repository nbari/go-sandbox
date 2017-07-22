package main

import (
	"log"
	"log/syslog"
)

func main() {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "emoji")
	if e == nil {
		log.SetOutput(logwriter)
	}

	log.Printf("emoji %c   %c\n", 0x2b55, 0x1f33f)
}
