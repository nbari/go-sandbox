package main

import "fmt"

func isPalindrome(str string) bool {
	var r string
	for _, s := range str {
		r = string(s) + r
	}
	return r == str
}
func main() {
	str := "hello"
	fmt.Println(isPalindrome(str))
	str = "madam"
	fmt.Println(isPalindrome(str))
	str = "ana"
	fmt.Println(isPalindrome(str))
}
