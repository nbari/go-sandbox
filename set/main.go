package main

import "fmt"

type item struct {
	key string
	val string
}

type Set []item

func (s *Set) Add(key, value string) {
	for k, v := range *s {
		if v.key == key {
			(*s)[k] = item{key, value}
			return
		}
	}
	item := item{key, value}
	*s = append(*s, item)
}

func (s *Set) Get(key string) (string, bool) {
	for _, v := range *s {
		if v.key == key {
			return v.val, true
		}
	}
	return "", false
}

func (s *Set) Delete(key string) {
	for k, v := range *s {
		if v.key == key {
			(*s)[k] = (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
			return
		}
	}
}

func main() {
	set := new(Set)
	set.Add("key", "value")
	set.Add("key", "value")
	set.Add("key", "value")
	set.Add("key", "value")
	set.Add("key2", "value")
	set.Add("key", "foo")
	fmt.Printf("set = %s\n", set)
	if val, ok := set.Get("key"); ok {
		fmt.Printf("val = %+v\n", val)
	}
	set.Add("key", "test")
	if val, ok := set.Get("key"); ok {
		fmt.Printf("val = %+v\n", val)
	}
	set.Delete("key")
	fmt.Printf("set = %s\n", set)
	set.Add("key", "mati")
	fmt.Printf("set = %s\n", set)
	set.Delete("key")
	fmt.Printf("set = %s\n", set)
	set.Delete("key2")
	fmt.Printf("set = %s\n", set)
	set.Add("1", "one")
	set.Add("2", "two")
	set.Add("3", "three")
	set.Add("4", "four")
	fmt.Printf("set = %s\n", set)
	set.Delete("2")
	fmt.Printf("set = %s\n", set)
}
