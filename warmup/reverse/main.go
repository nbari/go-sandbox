package main

import "fmt"

func main() {
	str := "Grettings!"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	var out string
	for _, s := range str {
		out = string(s) + out
	}
	println(out)

	println(Reverse(str))
}

// recursion
func Reverse(str string) string {
	if str != "" {
		fmt.Printf("str = %+v\n", str)
		return Reverse(str[1:]) + str[:1]
	}
	return ""
}
