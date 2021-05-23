package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandRunes(t *testing.T) {
	type data struct {
		in     runes
		result []byte
	}

	tests := []data{
		{runes{[]rune{'a', 'Z'}}, []byte("a ` _ ^ ] \\ [ Z")},
		{runes{[]rune{'a', 'e'}}, []byte("a b c d e")},
		{runes{[]rune{'😄', '😁'}}, []byte("😄 😃 😂 😁")},
	}

	for _, tt := range tests {
		result := bytes.Join(tt.in.expand(), space)
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}
