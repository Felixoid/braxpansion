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
	"dfi mon-mon{i,y,ie} {13..050}",
	"metric.{us,ru,en,de,dk,gb,in}server{01..12}.cpu.{0..3}.{idle,sys,user}",
}

var inBytes = stringsToBytesSlice(input)

func stringsToBytesSlice(in []string) [][]byte {
	out := make([][]byte, len(in))
	for i := range in {
		out[i] = []byte(in[i])
	}
	return out
}

func shBraceExpansion(in string) []string {
	p := syntax.NewParser()
	w, err := p.Document(strings.NewReader(in))
	if err != nil {
		return []string{}
	}
	words := expand.Braces(w)
	out := make([]string, 0, 1)
	b := new(strings.Builder)
	for _, w := range words {
		for _, p := range w.Parts {
			if lit, ok := p.(*syntax.Lit); ok {
				b.WriteString(lit.Value)
			}
		}
		out = append(out, b.String())
		b.Reset()
	}
	return out
}

func bytesAsString(in []byte) [][]byte {
	result := braxpansion.ExpandString(string(in))
	return stringsToBytesSlice(result)
}

func BenchmarkExpand(b *testing.B) {
	benchSize := []string{"Tiny", "Small", "Big", "Huge"}
	type benchType struct {
		name           string
		stringFunction func(string) []string
		bytesFunction  func([]byte) [][]byte
	}

	benchmarks := []benchType{
		{"DummyWarmUp", func(in string) []string { return []string{in} }, func(in []byte) [][]byte { return [][]byte{in} }},
		{"ExpandString", braxpansion.ExpandString, nil},
		{"ExpandBytes", nil, braxpansion.ExpandBytes},
		{"ExpBytesAsStr", nil, bytesAsString},
		{"Gobrex", gobrex.Expand, nil},
		{"ShExpand", shBraceExpansion, nil},
	}

	for s, name := range benchSize {
		s++
		for _, bench := range benchmarks {
			b.Run(name+"-"+bench.name, func(b *testing.B) {
				if bench.stringFunction != nil {
					for i := 0; i < b.N; i++ {
						for _, in := range input[:s] {
							bench.stringFunction(in)
						}
					}
				}
				if bench.bytesFunction != nil {
					for i := 0; i < b.N; i++ {
						for _, in := range inBytes[:s] {
							bench.bytesFunction(in)
						}
					}
				}
			})
		}
		b.Log("\n")
	}
}
