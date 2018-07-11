package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-yaml/yaml"
)

var data = `
sentinel:
  number: 3
server:
  number: 7
config:
  fere_size: 5
lcmea:
  eza_ze: all
`

func main() {
	m := make(map[string]map[string]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	out := []string{}
	for k, v := range m {
		for j, i := range v {
			out = append(out, fmt.Sprintf("%s.%v=%v", k, j, i))
		}
	}

	fmt.Println(strings.Join(out, ", "))
}
