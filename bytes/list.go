package bytes

import "bytes"

// list is a coma separated strings. It could contain nexted expression
type list struct {
	body []byte
}

func (l list) expand() ([][]byte, error) {
	getNext := func(b []byte) int {
		nextComa := bytes.IndexRune(b, ',')
		if nextComa == -1 {
			return len(b)
		}
		return nextComa
	}

	result := make([][]byte, 0, bytes.Count(l.body, coma))
	coma := 0
	for {
		nextComa := coma + getNext(l.body[coma:])
		brace := bytes.IndexRune(l.body[coma:nextComa], '{')
		if brace == -1 {
			result = append(result, l.body[coma:nextComa])
		} else {
			// brace inside the list could be only in pairs
			_, stop := getPair(l.body[coma:])
			for {
				brace = bytes.IndexRune(l.body[coma+stop:], '{')
				if brace != -1 {
					_, newStop := getPair(l.body[coma+stop+1:])
					stop += newStop + 1
					continue
				}
				nextComa = coma + stop + getNext(l.body[coma+stop:])
				break
			}
			expanded, err := expandSingle(l.body[coma:nextComa])
			result = append(result, expanded...)
			if err != nil {
				return result, err
			}
		}

		if nextComa == len(l.body) {
			break
		}

		coma = nextComa + 1
	}
	return result, nil
}
