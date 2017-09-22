package test_splitpath

import (
	"fmt"
	"strings"
	"testing"
)

// splitPath returns an slice of the path
func splitPath(b *testing.B, p string) {
	b.ReportAllocs()
	b.ResetTimer()
	//	var pathParts []string

	for i := 0; i < b.N; i++ {
		pathParts := strings.FieldsFunc(p, func(c rune) bool {
			return c == '/'
		})
		// root (empty slice)
		if len(pathParts) == 0 {
			pathParts = append(pathParts, "/")
		}
		p = fmt.Sprintf("%s/%d", p, i)
	}

}

func BenchmarkSplitPath(b *testing.B) {
	splitPath(b, "")
}
