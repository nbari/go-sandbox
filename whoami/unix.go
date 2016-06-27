package main

import (
	"os"
	"os/user"
	"strconv"
)

func getUser() string {
	uid := strconv.Itoa(os.Geteuid())
	u, err := user.LookupId(uid)
	if err != nil {
		fatal.Fatalf("cannot find name for user ID %s\n", uid)
	}
	return u.Username
}
