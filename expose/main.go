package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/nbari/violetear"
)

type Repos struct {
	Remotes []string
}

func findRemotes(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob("*")
	repos := &Repos{}
	for _, f := range files {
		fi, err := os.Stat(f)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if fi.Mode().IsDir() {
			out, err := exec.Command("git",
				fmt.Sprintf("--git-dir=%s/.git", f),
				"remote",
				"get-url",
				"--push",
				"--all",
				"origin").Output()
			if err != nil {
				log.Println(err)
				continue
			}
			if len(out) > 0 {
				repos.Remotes = append(repos.Remotes, fmt.Sprintf("%s", out))
			}
		}
	}
	if err := json.NewEncoder(w).Encode(repos); err != nil {
		log.Println(err)
	}
}

func main() {
	router := violetear.New()
	router.HandleFunc("*", findRemotes)
	log.Fatal(http.ListenAndServe(":8080", router))
}
