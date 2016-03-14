// sudo go run main.go -gid 20 -uid 501
// http://stackoverflow.com/questions/35962179/is-there-a-way-to-chroot-sandbox-a-go-os-exec-call-prevent-rm-rf
//
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var oUid = flag.Int("uid", 0, "Run with User ID")
	var oGid = flag.Int("gid", 0, "Run with Group ID")
	flag.Parse()

	// Get UID/GUID from args
	var uid = *oUid
	var gid = *oGid

	// Run whoami
	out, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Output whoami
	log.Println("Original UID/GID whoami:", string(out))
	log.Println("Setting UID/GUID")

	// Change privileges
	err = syscall.Setgid(gid)
	if err != nil {
		log.Println("Cannot setgid")
		log.Fatal(err)
		return
	}

	err = syscall.Setuid(uid)
	if err != nil {
		log.Println("Cannot setuid")
		log.Fatal(err)
		return
	}

	// Execute whoami again
	out, err = exec.Command("whoami").Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Changed UID/GID whoami:", string(out))

	// Do some dangerous stuff
	log.Println("Creating a executable file within /bin should fail...")
	_, err = os.Create("/bin/should-fail")
	if err == nil {
		log.Println("Warning: operation did not fail")
		return
	}

	log.Println("We are fine", err)
}
