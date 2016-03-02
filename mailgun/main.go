package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	client := &http.Client{}
	URL := "https://api.mailgun.net/v3/your-domain/messages"
	v := url.Values{}
	v.Set("from", "from@your-domain")
	v.Set("to", "to@domain")
	v.Set("subject", "subject")
	v.Set("text", "body")
	req, _ := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
	req.SetBasicAuth("api", "key-XXX")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body: %s", body)
}
