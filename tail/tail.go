package main

import (
	"fmt"
	"os"
	"time"
)

func watchFile(file string) error {
	s, err := os.Stat(file)
	if err != nil {
		return err
	}

	for {
		s, err := os.Stat(file)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("vim-go")
}
