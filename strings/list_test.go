package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandList(t *testing.T) {
	type data struct {
		in     list
		result string
	}
	tests := []data{
		{list{"a,b,c,d"}, "a b c d"},
		{list{"a,"}, "a "},
		{list{"abc,{ab,ce}{bc,de}"}, "abc abbc abde cebc cede"},
		{list{"{a,{b..e}{a..c}}"}, "a ba bb bc ca cb cc da db dc ea eb ec"},
	}

	for _, tt := range tests {
		result := strings.Join(tt.in.expand(), space)
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}
