package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type File struct {
	path string
	size int64
	uid  uint32
}

// A data structure to hold key/value pairs
type Pair struct {
	Key   uint32
	Value int64
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

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
	ctx, cf := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cf()
	m, err := search(ctx, path)
	if err != nil {
		log.Fatal(err)
	}

	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	for _, v := range p {
		u, err := user.LookupId(fmt.Sprintf("%d", v.Key))
		if err != nil {
			fmt.Printf("UID: %d: %d\n", v.Key, v.Value)
		} else {
			fmt.Printf("%s: %d\n", u.Username, v.Value)
		}
	}
}

func search(ctx context.Context, root string) (map[uint32]int64, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan File, 100)

	// get all the paths

	g.Go(func() error {
		defer close(paths)

		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- File{path, info.Size(), info.Sys().(*syscall.Stat_t).Uid}:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})

	})

	go func() {
		g.Wait()
	}()

	uids := make(map[uint32]int64)
	for r := range paths {
		if val, ok := uids[r.uid]; ok {
			uids[r.uid] = val + r.size
		} else {
			uids[r.uid] = r.size
		}
	}
	return uids, g.Wait()
}
