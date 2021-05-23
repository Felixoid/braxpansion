package bytes

import (
	"bytes"
)

var (
	coma  = []byte(",")
	dots  = []byte("..")
	space = []byte(" ")
)

// Expand takes the []byte contains the shell expansion expression and returns a slice of []byte after
// they are expanded. As in shells, each word is processed separately, so `12{1,2,3,4}as ds{1..3}22` produces `121as 122as 123as 124as ds122 ds222 ds322`
func Expand(in []byte) ([][]byte, error) {
	fields := bytes.Fields(in)
	result := make([][]byte, 0, len(fields))
	for _, f := range fields {
		expanded, err := expandSingle(f)
		result = append(result, expanded...)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

// expandSingle expands single field and multiplies all braces pairs in it with plain text and each other.
func expandSingle(in []byte) ([][]byte, error) {
	start, stop := getPair(in)
	cur := 0
	exps := make([][][]byte, 0)
	dimensions := make([]int, 0)
	resLen := 1

	for {
		if start == -1 {
			break
		}
		if start != 0 {
			exps = append(exps, [][]byte{in[cur : start+cur]})
			dimensions = append(dimensions, 1)
		}
		exp, err := getExpression(in[cur+start : cur+stop+1]).expand()
		exps = append(exps, exp)
		if err != nil {
			return [][]byte{}, err
		}
		dimensions = append(dimensions, len(exp))
		resLen *= len(exp)
		cur += stop + 1
		if cur == len(in) {
			break
		}
		start, stop = getPair(in[cur:])
	}

	if cur != len(in) {
		exps = append(exps, [][]byte{in[cur:]})
		dimensions = append(dimensions, 1)
	}

	result := make([][]byte, resLen)
	curMult := make([]int, len(exps))
	for i := 0; i < resLen; i++ {
		b := new(bytes.Buffer)
		b.Grow(len(in))
		for i := range exps {
			_, err := b.Write(exps[i][curMult[i]])
			if err != nil {
				return [][]byte{}, err
			}
		}
		for j := len(curMult) - 1; 0 <= j; j-- {
			curMult[j]++
			if curMult[j] == dimensions[j] {
				curMult[j] = 0
				continue
			}
			break
		}
		result[i] = b.Bytes()
	}
	return result, nil
}

func getPair(in []byte) (start, stop int) {
	start = bytes.IndexRune(in, '{')
	stop = bytes.IndexRune(in, '}')
	if start == -1 || stop == -1 || stop < start {
		return -1, -1
	}

	depth := 0
	cur := start
	for {
		unpair := bytes.IndexRune(in[cur+1:stop], '{')
		if unpair != -1 {
			cur += unpair + 1
			depth++
			continue
		}
		// the only {} pair
		if depth == 0 {
			break
		}

		pair := bytes.IndexRune(in[stop+1:], '}')
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
