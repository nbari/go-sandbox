package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w := os.Stdout
	for {
		r, err := os.Open("file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
