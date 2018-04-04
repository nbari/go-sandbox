package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func walkDir(dir string) map[uint32]int64 {
	defer wg.Done()

	uids := make(map[uint32]int64)
	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() && path != dir {
			wg.Add(1)
			go walkDir(path)
			return filepath.SkipDir
		}
		if f.Mode().IsRegular() {
			uid := f.Sys().(*syscall.Stat_t).Uid
			if val, ok := uids[uid]; ok {
				uids[uid] = val + f.Size()
			} else {
				uids[uid] = f.Size()
			}
		}
		return nil
	}

	filepath.Walk(dir, visit)
	return uids
}

func main() {
	var path string
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [dir]\n\n", os.Args[0])
		fmt.Printf("  dir   The directory that will be scanned.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) >= 1 {
		path = flag.Args()[0]
	} else {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()

	wg.Add(1)
	m := walkDir(path)
	wg.Wait()
	fmt.Printf("m = %+v\n", m)
}
