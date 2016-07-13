package main

import (
	"encoding/json"
	"github.com/nbari/violetear"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	//	ip := strings.Split(r.RemoteAddr, ":")[0]
	j := map[string]string{
		"IP":   r.RemoteAddr,
		"UA":   r.Header.Get("User-Agent"),
		"Time": time.Now().UTC().Format(time.RFC3339),
	}
	if err := json.NewEncoder(w).Encode(j); err != nil {
		panic(err)
	}
}

func viewHeaders(w http.ResponseWriter, r *http.Request) {
	j := make(map[string]interface{})
	for k, v := range r.Header {
		j[k] = v
	}
	if err := json.NewEncoder(w).Encode(j); err != nil {
		panic(err)
	}
}

func main() {
	router := violetear.New()
	router.HandleFunc("*", index)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
