package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandNumbers(t *testing.T) {
	type data struct {
		in     expression
		result string
	}

	tests := []data{
		{getExpression("{-10..1}"), "-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 0 1"},
		{getExpression("{1..5}"), "1 2 3 4 5"},
		{getExpression("{1..1}"), "1"},
		{getExpression("{1..005..2}"), "001 003 005"},
		{getExpression("{001..5..2}"), "001 003 005"},
		{getExpression("{1..005..-002}"), "0005 0003 0001"},
		{getExpression("{3..-02..3}"), "003 000"},
		{getExpression("{3..-02..-3}"), "000 003"},
		{getExpression("{3..-03..-3}"), "-03 000 003"},
		{getExpression("{3..-03..0}"), "{3..-03..0}"},
	}

	for _, tt := range tests {
		exp := tt.in
		assert.IsType(t, numbers{}, exp)
		result := strings.Join(tt.in.expand(), space)
		assert.Equal(t, tt.result, result, "input %q", tt.in)
	}
}

func TestNumbersPanics(t *testing.T) {
	tt := numbers{orig: []string{"-1", "-2", "1"}, seq: []int{-1, -2}}
	assert.Panics(t, func() { tt.expand() })
}
