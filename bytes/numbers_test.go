package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandNumbers(t *testing.T) {
	type data struct {
		in     expression
		result []byte
	}

	tests := []data{
		{getExpression([]byte("{-10..1}")), []byte("-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 0 1")},
		{getExpression([]byte("{1..5}")), []byte("1 2 3 4 5")},
		{getExpression([]byte("{1..005..2}")), []byte("001 003 005")},
		{getExpression([]byte("{1..005..-002}")), []byte("0005 0003 0001")},
		{getExpression([]byte("{3..-02..3}")), []byte("003 000")},
		{getExpression([]byte("{3..-02..-3}")), []byte("000 003")},
		{getExpression([]byte("{3..-03..-3}")), []byte("-03 000 003")},
	}

	for _, tt := range tests {
		exp := tt.in
		assert.IsType(t, numbers{}, exp)
		result := bytes.Join(tt.in.expand(), space)
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}
