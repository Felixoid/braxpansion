package strings

import (
	"strings"
)

const (
	coma  = ","
	dots  = ".."
	space = " "
)

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
		// Optimization for pre-defined buffer length to avoid reallocations
		b.Grow(len(in))
		for i := range exps {
			// Builder.WriteString always returns nil as error
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

// getPair returns the top level expression. If the first `{` doesn't have the pair,
// it recursively executed for the substring after it
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
