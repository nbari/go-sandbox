// http://www.sarathlakshman.com/2010/08/12/exit-from-chroot-environment-python/

package main

import (
	"fmt"
	"syscall"
)

func main() {

	err := syscall.Chroot("/tmp/root")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

}
