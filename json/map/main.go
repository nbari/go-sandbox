package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

const jsonStream = `
    {"xyz": {
        "name": "john",
        "age": 23,
        "xyz": "weu33s"}
    }`

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Info map[string]Person

func main() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var info Info
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d\n", info["xyz"].Name, info["xyz"].Age)
	}
}
