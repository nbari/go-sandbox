package main

import (
	"fmt"
	"sort"
)

type mySet struct {
	data map[string]int
}
type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (m *mySet) Add(item string) {
	for k, v := range m.data {
		if v == 0 {
			delete(m.data, k)
			continue
		}
		m.data[k] = v - 1
	}
	m.data[item] = 10
}

func main() {
	x := &mySet{
		data: make(map[string]int),
	}

	for i := 0; i < 10; i++ {
		x.data[fmt.Sprintf("test-%d", i)] = i
	}

	x.Add("foo")

	for i := 0; i < 3; i++ {
		x.Add(fmt.Sprintf("foo-%d", i))
	}

	pl := make(PairList, len(x.data))
	i := 0
	for k, v := range x.data {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	fmt.Printf("pl = %+v\n", pl)

}
