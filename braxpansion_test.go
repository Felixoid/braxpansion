package braxpansion

import (
	"strings"
	"testing"

	gobrex "github.com/kujtimiihoxha/go-brace-expansion"
	"github.com/stretchr/testify/assert"
	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/syntax"
)

func TestExpandList(t *testing.T) {
	type data struct {
		in     list
		result string
	}
	tests := []data{
		{list{"a,b,c,d"}, "a b c d"},
		{list{"abc,{ab,ce}{bc,de}"}, "abc abbc abde cebc cede"},
		{list{"{a,{b..e}{a..c}}"}, "a ba bb bc ca cb cc da db dc ea eb ec"},
	}

	for _, tt := range tests {
		result := strings.Join(tt.in.expand(), " ")
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}

func TestExpandNumbers(t *testing.T) {
	type data struct {
		in     expression
		result string
	}

	tests := []data{
		{getExpression("{-10..1}"), "-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 0 1"},
		{getExpression("{1..5}"), "1 2 3 4 5"},
		{getExpression("{1..005..2}"), "001 003 005"},
		{getExpression("{1..005..-002}"), "0005 0003 0001"},
		{getExpression("{3..-02..3}"), "003 000"},
		{getExpression("{3..-02..-3}"), "000 003"},
		{getExpression("{3..-03..-3}"), "-03 000 003"},
	}

	for _, tt := range tests {
		exp := tt.in
		assert.IsType(t, numbers{}, exp)
		result := strings.Join(tt.in.expand(), " ")
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}

func TestExpandRunes(t *testing.T) {
	type data struct {
		in     runes
		result string
	}

	tests := []data{
		{runes{[]rune{'a', 'Z'}}, "a ` _ ^ ] \\ [ Z"},
		{runes{[]rune{'a', 'e'}}, "a b c d e"},
		{runes{[]rune{'😄', '😁'}}, "😄 😃 😂 😁"},
	}

	for _, tt := range tests {
		result := strings.Join(tt.in.expand(), " ")
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}

func TestExpandSingle(t *testing.T) {
	type data struct {
		in  string
		out string
	}

	tests := []data{
		{in: "1{b..e}2{a..c}3", out: "1b2a3 1b2b3 1b2c3 1c2a3 1c2b3 1c2c3 1d2a3 1d2b3 1d2c3 1e2a3 1e2b3 1e2c3"},
		{in: "as{12,32}{a..c}{2}", out: "as12a{2} as12b{2} as12c{2} as32a{2} as32b{2} as32c{2}"},
	}

	for _, tt := range tests {
		result := strings.Join(expandSingle(tt.in), " ")
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}

func TestGetExpression(t *testing.T) {
	type data struct {
		in  string
		exp expression
	}

	tests := []data{
		{"{1,2,3,4}", list{"1,2,3,4"}},
		{"{1,2,3,4,{1,2,3}}", list{"1,2,3,4,{1,2,3}"}},
		{"{-12..3}", numbers{[]string{"-12", "3"}, []int{-12, 3}}},
		{"{1..5}", numbers{[]string{"1", "5"}, []int{1, 5}}},
		{"{1..005..2}", numbers{[]string{"1", "005", "2"}, []int{1, 5, 2}}},
		{"{1..a}", runes{seq: []rune{'1', 'a'}}},
		{"{😁..👌}", runes{seq: []rune{'😁', '👌'}}},
		{"{13323}", none{"{13323}"}},
	}

	for _, tt := range tests {
		exp := getExpression(tt.in)
		assert.Equal(t, tt.exp, exp, "body %s", tt.in)
	}
}

func TestGetPair(t *testing.T) {
	type data struct {
		in    string
		start int
		stop  int
	}
	tests := []data{
		{"{x{12,{}}xxxxx", 2, 8},
		{"x{12,{}}{{,13}", 1, 7},
		{"{x{12,{}}{{,13}", 2, 8},
		{"{{x{12,{}}{{,13}", 3, 9},
	}

	for _, tt := range tests {
		start, stop := getPair(tt.in)
		assert.Equal(t, tt.start, start, "start of %s", tt.in)
		assert.Equal(t, tt.stop, stop, "stop of %s", tt.in)
	}
}

var input = []string{
	"1{b..e}2{a..c}3",
	"232{ad,fdff,wwwww,asdasd{02..3}}{z..A}",
}

func BenchmarkExpand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, i := range input {
			Expand(i)
		}
	}
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
