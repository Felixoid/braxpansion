package braxpansion

import (
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
