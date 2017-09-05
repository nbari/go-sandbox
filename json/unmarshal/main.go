package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	i := Item{}
	if err := json.Unmarshal([]byte(jsonStream), &i); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	var (
		mem int
		cpu float64
	)
	for _, v := range i.Apps {
		cpu += v.Cpus
		mem += v.Mem
	}
	fmt.Printf("cpu = %.2f\n", cpu)
	fmt.Printf("mem = %+v\n", mem)
}
