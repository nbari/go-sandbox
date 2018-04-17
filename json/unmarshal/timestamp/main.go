package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const jsonStream = `[
  {
    "timeStamp": 1227232494001,
    "y0": "6.44",
    "y1": "-3.96",
    "y2": "-7.89",
    "y3": "4.26"
  },
  {
    "timeStamp": 1228472111570,
    "y0": "6.31",
    "y1": "-4.38",
    "y2": "-7.60",
    "y3": "3.65"
  }
]`

type Record struct {
	TimeStamp int64
	Y0        float64 `json:",string"`
	Y1        float64 `json:",string"`
	Y2        float64 `json:",string"`
	Y3        float64 `json:",string"`
}

func main() {
	i := []Record{}
	if err := json.Unmarshal([]byte(jsonStream), &i); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, v := range i {
		fmt.Printf("timestamp: %v y0: %0.2f y1: %0.2f y2: %0.2f y3: %0.2f\n",
			v.TimeStamp,
			v.Y0,
			v.Y1,
			v.Y2,
			v.Y3)
	}
}
