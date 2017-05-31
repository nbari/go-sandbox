package main

import (
	"fmt"
	"strings"
	"unicode"
)

const foo = `time="2017-05-30T19:02:08-05:00" level=info msg="some log message" app=sample size=10`

func main() {
	lastQuote := rune(0)
	f := func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return false
		default:
			return unicode.IsSpace(c)

		}
	}

	// splitting string by space but considering quoted section
	items := strings.FieldsFunc(foo, f)

	// create and fill the map
	m := make(map[string]string)
	for _, item := range items {
		x := strings.Split(item, "=")
		m[x[0]] = x[1]
	}

	// print the map
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}
