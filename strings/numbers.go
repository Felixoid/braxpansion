package strings

import "fmt"

// numbers generates sequense from first to second element. There is a possible third element `step` with sign
type numbers struct {
	orig []string
	seq  []int
}

func (n numbers) expand() ([]string, error) {
	if len(n.seq) != len(n.orig) {
		// This is some error
		return nil, fmt.Errorf("the length of original and parsed arguments aren't the same: %d != %d", len(n.orig), len(n.seq))
	}
	first := n.seq[0]
	last := n.seq[1]
	step := 1
	reversed := false
	if len(n.seq) == 3 {
		step = n.seq[2]
		if step == 0 {
			return []string{fmt.Sprintf("%s..%s..0", n.orig[0], n.orig[1])}, nil
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

	return result, nil
}
