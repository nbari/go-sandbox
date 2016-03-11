package main

import (
	"fmt"
	"os"
)

func main() {

	var shell = os.Getenv("SHELL")
	if "" != shell {
		fmt.Println(shell, "-c")
	}
}
