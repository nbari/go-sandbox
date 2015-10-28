package trie_slice

import (
	"net/http"
	"strconv"
	"testing"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

var t = NewTrie()

func Benchmark_Set(b *testing.B) {
	path := []string{"root"}
	for i := 0; i < b.N; i++ {
		si := strconv.Itoa(i)
		path = append(path, si)
		t.Set(path, messageHandler(si), "ALL")
		//		fmt.Println(path)
	}
}

func Benchmark_Get(b *testing.B) {
	path := []string{"root"}
	for i := 0; i < b.N; i++ {
		si := strconv.Itoa(i)
		path = append(path, si)
		t.Get(path)
	}
}
