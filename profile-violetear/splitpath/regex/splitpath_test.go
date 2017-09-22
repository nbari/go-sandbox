package test_splitpath

import (
	"fmt"
	"regexp"
	"testing"
)

var splitPathRx = regexp.MustCompile(`[^/ ]+`)

// splitPath returns an slice of the path
func splitPath(b *testing.B, p string) {
	b.ReportAllocs()
	b.ResetTimer()
	//	var pathParts []string

	for i := 0; i < b.N; i++ {
		//for i := 0; i < 10; i++ {
		pathParts := splitPathRx.FindAllString(p, -1)
		// root (empty slice)
		if len(pathParts) == 0 {
			pathParts = append(pathParts, "/")
		}
		p = fmt.Sprintf("%s/%d", p, i)
		//	fmt.Printf("p = %+v\n", p)
		//	fmt.Printf("pathParts = %+v\n", pathParts)
	}

}

func BenchmarkSplitPath(b *testing.B) {
	splitPath(b, "")
}
