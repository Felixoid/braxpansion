package bench

import (
	"strings"
	"testing"

	"github.com/Felixoid/braxpansion"
	gobrex "github.com/kujtimiihoxha/go-brace-expansion"
	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/syntax"
)

var input = []string{
	"1{b..e}2{a..c}3",
	"232{ad,fdff,wwwww,asdasd{02..3}}{z..A}",
}

var inByte = [][]byte{
	[]byte("1{b..e}2{a..c}3"),
	[]byte("232{ad,fdff,wwwww,asdasd{02..3}}{z..A}"),
}

func BenchmarkGobrex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, i := range input {
			gobrex.Expand(i)
		}
	}
}

func BenchmarkShExpand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, i := range input {
			p := syntax.NewParser()
			w, err := p.Document(strings.NewReader(i))
			if err != nil {
				b.Fatal(err.Error())
			}
			expand.Braces(w)
		}
	}
}

func BenchmarkExpandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, i := range input {
			braxpansion.ExpandString(i)
		}
	}
}

func BenchmarkExpandBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, i := range inByte {
			braxpansion.ExpandBytes(i)
		}
	}
}
func BenchmarkExpandBytesAsString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, i := range inByte {
			str := string(i)
			result := braxpansion.ExpandString(str)
			for _, r := range result {
				_ = string(r)
			}
		}
	}
}
