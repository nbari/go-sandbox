package main

import (
	"fmt"
	"os"
	"testing"
)

var exitCode int

func Test_main(t *testing.T) {
	fmt.Println("aqui test_____main")
	go main()
	exitCode = <-exitCh
}

func TestMain(m *testing.M) {
	fmt.Println("maiiiiiii")
	m.Run()
	// can exit because cover profile is already written
	os.Exit(exitCode)
}
