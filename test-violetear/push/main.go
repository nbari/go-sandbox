package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

const mainJS = `console.log("hello world");`

const indexHTML = `<html>
<head>
	<title>Hello</title>
	<script src="/main.js"></script>
</head>
<body>
</body>
</html>
`

func rootJS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, mainJS)
}

func root(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push("/main.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	fmt.Fprintf(w, indexHTML)
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("/", root)
	router.HandleFunc("main.js", rootJS)

	log.Fatal(http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", router))
}
