package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nbari/go-sandbox/decorator/client"
)

func main() {
	// c := client.Decorate(&http.Client{},
	c := client.Decorate(client.MockClient{},
		client.Last(log.New(os.Stdout, "[last decorator] ", 0)),
		client.Logging(log.New(os.Stdout, "[logging decorator] ", 0)),
		client.Audit(log.New(os.Stdout, "[audit decorator] ", 0)),
		client.First(log.New(os.Stdout, "[first decorator] ", 0)),
	)
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	c.Do(req)
}
