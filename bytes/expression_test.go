package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExpression(t *testing.T) {
	type data struct {
		in  []byte
		exp expression
	}

	tests := []data{
		{[]byte("{1,2,3,4}"), list{[]byte("1,2,3,4")}},
		{[]byte("{1,2,3,4,{1,2,3}}"), list{[]byte("1,2,3,4,{1,2,3}")}},
		{[]byte("{-12..3}"), numbers{[][]byte{[]byte("-12"), []byte("3")}, []int{-12, 3}}},
		{[]byte("{1..5}"), numbers{[][]byte{[]byte("1"), []byte("5")}, []int{1, 5}}},
		{[]byte("{1..005..2}"), numbers{[][]byte{[]byte("1"), []byte("005"), []byte("2")}, []int{1, 5, 2}}},
		{[]byte("{1..a}"), runes{seq: []rune{'1', 'a'}}},
		{[]byte("{😁..👌}"), runes{seq: []rune{'😁', '👌'}}},
		{[]byte("{1..s..w}"), none{[]byte("{1..s..w}")}},
		{[]byte("{1..as}"), none{[]byte("{1..as}")}},
		{[]byte("{13323}"), none{[]byte("{13323}")}},
	}

	for _, tt := range tests {
		exp := getExpression(tt.in)
		assert.Equal(t, tt.exp, exp, "body %s", tt.in)
	}
}
