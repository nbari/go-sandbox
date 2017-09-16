package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(io.LimitReader(response.Body, 40000))
	fmt.Printf("data = %sn", data)
}
