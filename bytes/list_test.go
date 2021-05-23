package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandList(t *testing.T) {
	type data struct {
		in     list
		result string
	}
	tests := []data{
		{list{[]byte("a,b,c,d")}, "a b c d"},
		{list{[]byte("a,")}, "a "},
		{list{[]byte("abc,{ab,ce}{bc,de}")}, "abc abbc abde cebc cede"},
		{list{[]byte("{a,{b..e}{a..c}}")}, "a ba bb bc ca cb cc da db dc ea eb ec"},
	}

	for _, tt := range tests {
		result, err := tt.in.expand()
		assert.NoError(t, err)
		assert.Equal(t, tt.result, string(bytes.Join(result, space)), "input %q", tt.in)
	}
}
