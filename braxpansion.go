// Package braxpansion provides shell-like braces expansion.
//
// Examples:
// list of coma-separated arguments `v{aa,bb,cc}e` => `vaae vbbe vcce`
// single rune range `{a..f}` => `a b c d e f`
// numbers range `1{9..11}` => `19 110 111`
// numbers range with leading zeros `2{9..011}` => `2009 2010 2011` OR `2{09..11}` => `209 210 211`
// numbers range with increment
package braxpansion

import (
	"fmt"
	"strconv"
	"strings"
)

// expression represents all possible expandable types
type expression interface {
	expand() []string
}

// none is not expandable and returns itself
type none struct {
	body string
}

func (n none) expand() []string {
	return []string{n.body}
}

// list is a coma separated strings. It could contain nexted expression
type list struct {
	body string
}

func (l list) expand() []string {
	getNext := func(b string) int {
		nextComa := strings.IndexRune(b, ',')
		if nextComa == -1 {
			return len(b)
		}
		return nextComa
	}

	result := make([]string, 0, strings.Count(l.body, ","))
	coma := 0
	for {
		nextComa := coma + getNext(l.body[coma:])
		brace := strings.IndexRune(l.body[coma:nextComa], '{')
		if brace == -1 {
			result = append(result, l.body[coma:nextComa])
		} else {
			// brace inside the list could be only in pairs
			_, stop := getPair(l.body[coma:])
			for {
				brace = strings.IndexRune(l.body[coma+stop:], '{')
				if brace != -1 {
					_, newStop := getPair(l.body[coma+stop+1:])
					stop += newStop + 1
					continue
				}
				nextComa = coma + stop + getNext(l.body[coma+stop:])
				break
			}
			expanded := expandSingle(l.body[coma:nextComa])
			result = append(result, expanded...)
		}

		if nextComa == len(l.body) {
			break
		}

		coma = nextComa + 1
	}
	return result
}

// numbers generates sequense from first to second element. There is a possible third element `step` with sign
type numbers struct {
	orig []string
	seq  []int
}

func (n numbers) expand() []string {
	if len(n.seq) != len(n.orig) {
		// This is some error
		return []string{"the unexpected error, please report it"}
	}
	first := n.seq[0]
	last := n.seq[1]
	step := 1
	reversed := false
	if len(n.seq) == 3 {
		step = n.seq[2]
		if step == 0 {
			return []string{fmt.Sprintf("%s..%s..0", n.orig[0], n.orig[1])}
		}
		if step < 0 {
			reversed = true
		}
	}

	if (last < first && 0 < step) || (first < last && step < 0) {
		step = -step
	}

	maxLen := 1
	for i := range n.seq {
		var add int
		if n.seq[i] < 0 {
			add = 1
		}
		curLen := add + 1
		curVal := n.seq[i]
		for {
			curVal /= 10
			if curVal == 0 {
				break
			}
			curLen++
		}
		needsAdjust := curLen < len(n.orig[i])
		curLen = len(n.orig[i])
		if maxLen < curLen && needsAdjust {
			maxLen = curLen
		}
	}

	result := make([]string, 0)
	for i := first; (first <= i && i <= last) || (last <= i && i <= first); i += step {
		result = append(result, fmt.Sprintf("%0*d", maxLen, i))
	}

	if reversed {
		for i := 0; i < len(result)/2; i++ {
			result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
		}
	}

	return result
}

// runes generates a sequence of runes from first to second element
type runes struct {
	seq []rune
}

func (r runes) expand() []string {
	first := int(r.seq[0])
	second := int(r.seq[1])
	step := 1
	l := second - first
	if second < first {
		step = -step
		l = -l
	}

	result := make([]string, 0, l)
	for i := first; i != second; i += step {
		result = append(result, string(rune(i)))
	}
	result = append(result, string(rune(second)))

	return result
}

// Expand takes the string contains the shell expansion expression and returns list of strings after
// they are expanded. As in shells, each word is processed separately, so `12{1,2,3,4}as ds{1..3}22` produces `121as 122as 123as 124as ds122 ds222 ds322`
func Expand(in string) []string {
	fields := strings.Fields(in)
	result := make([]string, 0, len(fields))
	for _, f := range fields {
		result = append(result, expandSingle(f)...)
	}
	return result
}

// expandSingle expands single field and multiplies all braces pairs in it with plain text and each other.
func expandSingle(in string) []string {
	start, stop := getPair(in)
	cur := 0
	exps := make([][]string, 0)
	dimensions := make([]int, 0)
	resLen := 1

	for {
		if start == -1 {
			break
		}
		if start != 0 {
			exps = append(exps, []string{in[cur : start+cur]})
			dimensions = append(dimensions, 1)
		}
		exp := getExpression(in[cur+start : cur+stop+1]).expand()
		exps = append(exps, exp)
		dimensions = append(dimensions, len(exp))
		resLen *= len(exp)
		cur += stop + 1
		if cur == len(in) {
			break
		}
		start, stop = getPair(in[cur:])
	}

	if cur != len(in) {
		exps = append(exps, []string{in[cur:]})
		dimensions = append(dimensions, 1)
	}

	result := make([]string, resLen)
	curMult := make([]int, len(exps))
	for i := 0; i < resLen; i++ {
		b := new(strings.Builder)
		for i := range exps {
			b.WriteString(exps[i][curMult[i]])
		}
		for j := len(curMult) - 1; 0 <= j; j-- {
			curMult[j]++
			if curMult[j] == dimensions[j] {
				curMult[j] = 0
				continue
			}
			break
		}
		result[i] = b.String()
	}
	return result
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

func getPair(in string) (start, stop int) {
	start = strings.IndexRune(in, '{')
	stop = strings.IndexRune(in, '}')
	if start == -1 || stop == -1 || stop < start {
		return -1, -1
	}

	depth := 0
	cur := start
	for {
		unpair := strings.IndexRune(in[cur+1:stop], '{')
		if unpair != -1 {
			cur += unpair + 1
			depth++
			continue
		}
		// the only {} pair
		if depth == 0 {
			break
		}

		pair := strings.IndexRune(in[stop+1:], '}')
		if pair == -1 {
			break
		}
		stop += pair + 1
		depth--
	}

	if depth != 0 {
		diff := start + 1
		subStart, subStop := getPair(in[diff:])
		return subStart + diff, subStop + diff
	}

	return
}
