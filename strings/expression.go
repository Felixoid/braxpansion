package strings

import (
	"strconv"
	"strings"
)

// expression represents all possible expandable types
type expression interface {
	expand() []string
}

// getExpression returns the top level expression. If the first `{` doesn't have the pair,
// it recursively executed for the substring after it
func getExpression(in string) expression {
	orig := in
	in = in[1 : len(in)-1]
	// Even if {,..g} is used, it's interpreted as a list
	if strings.IndexRune(in, ',') != -1 {
		return list{in}
	}

	if strings.Index(in, "..") == -1 {
		return none{orig}
	}

	args := strings.Split(in, "..")
	if len(args) != 2 && len(args) != 3 {
		return none{orig}
	}

	isNumbers := true
	nOrig := make([]string, len(args))
	nSeq := make([]int, len(args))
	for i, a := range args {
		n, err := strconv.Atoi(a)
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
		r := []rune(a)
		if len(r) != 1 {
			return none{orig}
		}
		rSeq[i] = r[0]
	}

	return runes{rSeq}
}
