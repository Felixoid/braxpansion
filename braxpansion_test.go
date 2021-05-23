package braxpansion

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandString(t *testing.T) {
	type data struct {
		in  string
		out string
	}

	tests := []data{
		{"a{b c}d", "a{b c}d"},
		{"a{b,c}d", "abd acd"},
		{"a{b,c}d a{b c}d", "abd acd a{b c}d"},
	}

	for _, tt := range tests {
		result := strings.Join(ExpandString(tt.in), " ")
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}

func TestExpandBytes(t *testing.T) {
	type data struct {
		in  []byte
		out []byte
	}

	tests := []data{
		{[]byte("a{b c}d"), []byte("a{b c}d")},
		{[]byte("a{b,c}d"), []byte("abd acd")},
		{[]byte("a{b,c}d a{b c}d"), []byte("abd acd a{b c}d")},
	}

	for _, tt := range tests {
		result := bytes.Join(ExpandBytes(tt.in), []byte(" "))
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}
