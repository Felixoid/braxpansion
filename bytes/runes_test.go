package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandRunes(t *testing.T) {
	type data struct {
		in     runes
		result string
	}

	tests := []data{
		{runes{[]rune{'a', 'Z'}}, "a ` _ ^ ] \\ [ Z"},
		{runes{[]rune{'a', 'e'}}, "a b c d e"},
		{runes{[]rune{'😄', '😁'}}, "😄 😃 😂 😁"},
	}

	for _, tt := range tests {
		result, err := tt.in.expand()
		assert.NoError(t, err)
		assert.Equal(t, tt.result, string(bytes.Join(result, space)), "input %q", tt.in)
	}
}
