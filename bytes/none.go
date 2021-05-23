package bytes

// none is not expandable and returns itself
type none struct {
	body []byte
}

func (n none) expand() ([][]byte, error) {
	return [][]byte{n.body}, nil
}
