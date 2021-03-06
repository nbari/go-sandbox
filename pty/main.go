package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	//c := exec.Command("grep", "--color=auto", "bar")
	c := exec.Command("sleep", "3")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4}) // EOT
	}()
	io.Copy(os.Stdout, f)
}
