package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandList(t *testing.T) {
	type data struct {
		in     list
		result []byte
	}
	tests := []data{
		{list{[]byte("a,b,c,d")}, []byte("a b c d")},
		{list{[]byte("a,")}, []byte("a ")},
		{list{[]byte("abc,{ab,ce}{bc,de}")}, []byte("abc abbc abde cebc cede")},
		{list{[]byte("{a,{b..e}{a..c}}")}, []byte("a ba bb bc ca cb cc da db dc ea eb ec")},
	}

	for _, tt := range tests {
		result := bytes.Join(tt.in.expand(), space)
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}
