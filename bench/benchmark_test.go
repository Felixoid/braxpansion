package bench

import (
	"strings"
	"testing"

	"github.com/Felixoid/braxpansion"
	gobrex "github.com/kujtimiihoxha/go-brace-expansion"
	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/syntax"
)

var input = map[string]string{
	"Tiny":  "1{b..e}2{a..c}3",
	"Small": "232{ad,fdff,wwwww,asdasd{02..3}}{Z..a}",
	"Big":   "dfi mon-mon{i,y,ie{a..c}}{13..050}",
	"Huge":  "metric.{us,ru,en,de,dk,gb,in}server{01..12}.cpu.{0..3}.{idle,sys,user}",
}

var inBytes = make(map[string][]byte, len(input))

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
	benchName := []string{"Tiny", "Small", "Big", "Huge"}
	for name, expr := range input {
		inBytes[name] = []byte(expr)
	}

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

	for _, name := range benchName {
		for _, bench := range benchmarks {
			b.Run(name+"-"+bench.name, func(b *testing.B) {
				if bench.stringFunction != nil {
					for i := 0; i < b.N; i++ {
						bench.stringFunction(input[name])
					}
				}
				if bench.bytesFunction != nil {
					for i := 0; i < b.N; i++ {
						bench.bytesFunction(inBytes[name])
					}
				}
			})
		}
		b.Log("\n")
	}
}
