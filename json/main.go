package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const jsonStream = `{
    "apps": [{
        "id": 1,
        "cpus": 0.1,
        "mem": 1
    }, {
        "id": 2,
        "cpus": 0.2,
        "mem": 2
    }, {
        "id": 3,
        "cpus": 0.3,
        "mem": 3
    }]
}`

type Item struct {
	Apps []Metrics
}

type Metrics struct {
	Id   int
	Cpus float64
	Mem  int
}

func main() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	i := Item{}
	for {
		if err := dec.Decode(&i); err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	var (
		mem int
		cpu float64
	)
	for _, v := range i.Apps {
		cpu += v.Cpus
		mem += v.Mem
	}
	fmt.Printf("cpu = %+v\n", cpu)
	fmt.Printf("mem = %+v\n", mem)
}
