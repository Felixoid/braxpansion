package strings

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
