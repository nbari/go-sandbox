package main

import (
	"os"
	"testing"
)

var exitCode int

func Test_main(t *testing.T) {
	go main()
	exitCode = <-exitCh
}

func TestMain(m *testing.M) {
	m.Run()
	// can exit because cover profile is already written
	os.Exit(exitCode)
}
