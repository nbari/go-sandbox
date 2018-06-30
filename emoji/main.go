package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Flag(x string) (string, error) {
	x = strings.ToUpper(x)
	if len(x) != 2 {
		return "", errors.New("country code must be two letters")
	}
	if x[0] < 'A' || x[0] > 'Z' || x[1] < 'A' || x[1] > 'Z' {
		return "", errors.New("invalid country code")
	}
	return string(0x1F1E6+rune(x[0])-'A') + string(0x1F1E6+rune(x[1])-'A'), nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Enter a country code")
		os.Exit(1)
	}
	f, err := Flag(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Flag: %s\n", f)
}
