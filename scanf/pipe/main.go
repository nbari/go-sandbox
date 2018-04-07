package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// check if there is somethinig to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("stdin = %s\n", stdin)
	} else {
		fmt.Println("Enter your name")

		var name string
		fmt.Scanf("%s", &name)
		fmt.Printf("name = %s\n", name)
	}
}
