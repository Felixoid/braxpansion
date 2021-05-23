package bytes

import (
	"bytes"
	"strconv"
)

// expression represents all possible expandable types
type expression interface {
	expand() ([][]byte, error)
}

// getExpression returns the top level expression. If the first `{` doesn't have the pair,
// it recursively executed for the substring after it
func getExpression(in []byte) expression {
	orig := in
	in = in[1 : len(in)-1]
	// Even if {,..g} is used, it's interpreted as a list
	if bytes.IndexRune(in, ',') != -1 {
		return list{in}
	}

	if bytes.Index(in, []byte("..")) == -1 {
		return none{orig}
	}

	args := bytes.Split(in, dots)
	if len(args) != 2 && len(args) != 3 {
		return none{orig}
	}

	isNumbers := true
	nOrig := make([][]byte, len(args))
	nSeq := make([]int, len(args))
	for i, a := range args {
		n, err := strconv.Atoi(string(a))
		if err != nil {
			isNumbers = false
			break
		}
		nOrig[i] = a
		nSeq[i] = n
	}
	if isNumbers {
		return numbers{nOrig, nSeq}
	}

	if len(args) != 2 {
		return none{orig}
	}
	rSeq := make([]rune, len(args))
	for i, a := range args {
		if len(a) != 1 {
			return none{orig}
		}
		rSeq[i] = rune(a[0])
	}

	return runes{rSeq}
}
