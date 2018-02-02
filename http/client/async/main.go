package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func AsyncGet(u []string, ch chan<- string) {
	for _, v := range u {
		go func(url string) {
			res, err := HTTPGet(url)
			if err != nil {
				ch <- fmt.Sprintf("%s, Error: %s", url, err)
				return
			}
			res.Body.Close()
			// do something with the request
			ch <- fmt.Sprintf("%s - ok", url)
		}(v)
	}
}

func HTTPGet(url string) (*http.Response, error) {
	timeout := 100

	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   60 * time.Second,
		ResponseHeaderTimeout: time.Duration(timeout) * time.Millisecond,
	}

	client := &http.Client{}
	client.Transport = tr

	// create a new request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "foo")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	urls := []string{
		"https://stackoverflow.com",
		"https://google.com",
		"https://example.com",
	}

	status := make(chan string)
	AsyncGet(urls, status)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-status)
	}
}
