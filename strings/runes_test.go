package strings

import (
	"strings"
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
		result := strings.Join(tt.in.expand(), " ")
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}
