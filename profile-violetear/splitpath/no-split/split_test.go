package test_split

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// splitPath returns an slice of the path
func splitPath(b *testing.B, p string) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pathParts := strings.FieldsFunc(p, func(c rune) bool {
			return c == '/'
		})
		// root (empty slice)
		if len(pathParts) == 0 {
			pathParts = append(pathParts, "/")
		}
		//		fmt.Printf("len(pathParts) = %+v\n", len(pathParts))
		p = fmt.Sprintf("%s/%d", p, i)
	}
}

func splitPath2(b *testing.B, p string) {
	b.ReportAllocs()
	b.ResetTimer()
	var (
		fields int
		field  bytes.Buffer
	)
	for i := 0; i < b.N; i++ {
		for _, rune := range p {
			if rune == '/' {
				if field.Len() > 0 {
					//					fmt.Printf("field = %+v\n", field.String())
					field.Reset()
					fields++
				}
				continue
			}
			field.WriteRune(rune)
		}
		if field.Len() > 0 {
			fields++
		}
		//		fmt.Printf("fields = %+v\n", fields)
		p = fmt.Sprintf("%s/%d", p, i)
	}
}

func BenchmarkSplitPath(b *testing.B) {
	splitPath(b, "")
}

func BenchmarkNoStrings(b *testing.B) {
	splitPath2(b, "")
}
