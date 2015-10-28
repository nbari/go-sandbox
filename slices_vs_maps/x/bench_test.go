package main

import (
	//	"fmt"
	"testing"
)

type Methods struct {
	Method  string
	Handler string
}

var methods = []string{"OPTIONS", "GET", "HEAD", "PATCH", "POST", "PUT", "DELETE", "TRACE", "CONNECT"}

var the_map = map[string]string{
	"OPTIONS": "1",
	"GET":     "1",
	"HEAD":    "1",
	"PATCH":   "1",
	"POST":    "1",
	"PUT":     "1",
	"DELETE":  "1",
	"TRACE":   "1",
	"CONNECT": "1",
}

var h = []Methods{
	{"OPTIONS", "1"},
	{"GET", "1"},
	{"HEAD", "1"},
	{"PATCH", "1"},
	{"POST", "1"},
	{"PUT", "1"},
	{"DELETE", "1"},
	{"TRACE", "1"},
	{"CONNECT", "1"},
}

var x = []Methods{}

// uses a slice
func BenchmarkSearch_SLICE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range h {
			if v.Method == "/ja" {
				break
			}
		}
	}
}

// uses a map
func BenchmarkSearch_MAP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := the_map["/ja"]; ok {
			continue
		}
	}
}
