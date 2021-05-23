package bytes

// runes generates a sequence of runes from first to second element
type runes struct {
	seq []rune
}

func (r runes) expand() [][]byte {
	first := int(r.seq[0])
	second := int(r.seq[1])
	step := 1
	l := second - first
	if second < first {
		step = -step
		l = -l
	}

	result := make([][]byte, 0, l)
	for i := first; i != second; i += step {
		result = append(result, []byte(string(rune(i))))
	}
	result = append(result, []byte(string(rune(second))))

	return result
}
