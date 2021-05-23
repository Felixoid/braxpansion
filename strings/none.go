package strings

// none is not expandable and returns itself
type none struct {
	body string
}

func (n none) expand() ([]string, error) {
	return []string{n.body}, nil
}
