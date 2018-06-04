package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var (
		err      error
		response *http.Response
		retries  int = 3
	)
	for retries > 0 {
		// response, err = http.Get("https://non-existen")
		response, err = http.Get("https://google.com/robots.txt")
		if err != nil {
			log.Println(err)
			retries -= 1
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("data = %s\n", data)
	}
}
