package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	bytesSent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Sent",
		Help: "rsync bytes sent",
	})
	bytesReceived = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Received",
		Help: "rsync bytes received",
	})
)

// isFile return true if path is a regular file
func isFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	if m := f.Mode(); !m.IsDir() && m.IsRegular() && m&400 != 0 {
		return true
	}
	return false
}

// tail read last lines of file
func tail(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	offset, err := file.Seek(0, io.SeekEnd)
	buffer := make([]byte, 1024, 1024)
	for {
		readBytes, err := file.ReadAt(buffer, offset)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
		offset += int64(readBytes)
		if readBytes != 0 {
			s := strings.Fields(string(buffer[:readBytes]))
			// 2017/09/01 03:10:25 [16720] sent 6075891 bytes  received 38779 bytes  total size 302215840
			//                              4                      7
			if len(s) > 7 {
				tx, _ := strconv.Atoi(s[4])
				rx, _ := strconv.Atoi(s[7])
				bytesSent.Set(float64(tx))
				bytesReceived.Set(float64(rx))
			}
		}
		time.Sleep(time.Second)
	}
}

func main() {
	f := flag.String("f", "", "log `file`")
	p := flag.Int("p", 8080, "Listen on `port`")
	flag.Parse()

	// check if file exist
	if *f == "" || !isFile(*f) {
		fmt.Fprintf(os.Stderr, "Missing log file, use (\"%s -h\") for help.\n", os.Args[0])
		os.Exit(1)
	}

	// every second read the log file
	go tail(*f)

	// expose metrics
	prometheus.MustRegister(bytesSent)
	prometheus.MustRegister(bytesReceived)
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", *p), nil)
}
