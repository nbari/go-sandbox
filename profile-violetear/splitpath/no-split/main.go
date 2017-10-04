// package  main
package test_split

import (
	"bytes"
	"fmt"
)

func noSplit(path string) {
	var (
		fields int
		field  bytes.Buffer
	)
	for _, rune := range path {
		if rune == '/' {
			if field.Len() > 0 {
				//				fmt.Printf("field = %+v\n", field.String())
				field.Reset()
				fields++
			}
			continue
		}
		field.WriteRune(rune)
	}
	if field.Len() > 0 {
		//		fmt.Printf("field = %+v\n", field.String())
		fields++
		//fmt.Printf("field = %+v\n", field.String())
	}
	fmt.Printf("fields = %+v\n", fields)
}

func main() {
	path := ""
	for i := 0; i < 10000; i++ {
		noSplit(path)
		path = fmt.Sprintf("%s/%d", path, i)
	}
}
