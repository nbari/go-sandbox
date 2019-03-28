package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var (
		w string
		s int
	)
	for scanner.Scan() {
		if size := len(scanner.Text()); size > s {
			w = scanner.Text()
			s = size
		}
	}
	fmt.Printf("w = %+v\n", w)
	fmt.Printf("s = %+v\n", s)
}
