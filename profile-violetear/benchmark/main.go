package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/nbari/violetear"
)

var (
	port      string
	sleepTime time.Duration
)

func main() {
	p := flag.Int("p", 8080, "Listen on `port`")
	s := flag.Int("s", 0, "Sleep time in `milliseconds`")
	t := flag.Int("t", 10, "Sampling poing in `seconds`")
	flag.Parse()
	port = fmt.Sprintf(":%d", *p)
	sleepTime = time.Duration(*s) * time.Millisecond
	samplingPointDuration := time.Duration(*t) * time.Second

	go func() {
		time.Sleep(samplingPointDuration)
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		runtime.ReadMemStats(&mem)
		var u uint64 = 1024 * 1024
		fmt.Printf("TotalAlloc,Alloc,HeapAlloc,HeapSys\n%d,%d,%d,%d\n",
			mem.TotalAlloc/u,
			mem.Alloc/u,
			mem.HeapAlloc/u,
			mem.HeapSys/u)
	}()

	startVioletear()
}

func hello(w http.ResponseWriter, r *http.Request) {
	if sleepTime > 0 {
		time.Sleep(sleepTime)
	} else {
		runtime.Gosched()
	}
	w.Write([]byte("hello"))
}

func oneWord(w http.ResponseWriter, r *http.Request) {
	if sleepTime > 0 {
		time.Sleep(sleepTime)
	} else {
		runtime.Gosched()
	}
	w.Write([]byte("1 word"))
}

func twoWords(w http.ResponseWriter, r *http.Request) {
	if sleepTime > 0 {
		time.Sleep(sleepTime)
	} else {
		runtime.Gosched()
	}
	w.Write([]byte("2 words"))
}

func manyWords(w http.ResponseWriter, r *http.Request) {
	if sleepTime > 0 {
		time.Sleep(sleepTime)
	} else {
		runtime.Gosched()
	}
	w.Write([]byte("many words"))
}

func startVioletear() {
	r := violetear.New()
	//r.LogRequests = true
	r.AddRegex(":word", `^\w+$`)
	r.HandleFunc("/", hello)
	r.HandleFunc("/:word", oneWord)
	r.HandleFunc("/:word/:word", twoWords)
	r.HandleFunc("/:word/:word/*", manyWords)
	http.ListenAndServe(port, r)
}
