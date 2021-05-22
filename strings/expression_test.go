package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExpression(t *testing.T) {
	type data struct {
		in  string
		exp expression
	}

	tests := []data{
		{"{1,2,3,4}", list{"1,2,3,4"}},
		{"{1,2,3,4,{1,2,3}}", list{"1,2,3,4,{1,2,3}"}},
		{"{-12..3}", numbers{[]string{"-12", "3"}, []int{-12, 3}}},
		{"{1..5}", numbers{[]string{"1", "5"}, []int{1, 5}}},
		{"{1..005..2}", numbers{[]string{"1", "005", "2"}, []int{1, 5, 2}}},
		{"{1..a}", runes{seq: []rune{'1', 'a'}}},
		{"{😁..👌}", runes{seq: []rune{'😁', '👌'}}},
		{"{13323}", none{"{13323}"}},
	}

	for _, tt := range tests {
		exp := getExpression(tt.in)
		assert.Equal(t, tt.exp, exp, "body %s", tt.in)
	}
}
