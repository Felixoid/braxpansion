package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	type data struct {
		in  string
		out string
	}

	tests := []data{
		{"a{b c}d", "a{b c}d"},
		{"a{b,c}d", "abd acd"},
		{"a{b,c}d a{b c}d", "abd acd a{b c}d"},
		{"a{{b,c}d a{b c}d", "a{bd a{cd a{b c}d"},
		{"a{{b,c}}d a{b c}d", "a{b}d a{c}d a{b c}d"},
	}

	for _, tt := range tests {
		result := strings.Join(Expand(tt.in), space)
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}

func TestExpandSingle(t *testing.T) {
	type data struct {
		in  string
		out string
	}

	tests := []data{
		{"{2..3}", "2 3"},
		{"a{{2..3}}b", "a{2}b a{3}b"},
		{"a{{{2..3}}}b", "a{{2}}b a{{3}}b"},
		{"a{{2}}b", "a{{2}}b"},
		{"a{{{2}}}b", "a{{{2}}}b"},
		{"a{{1,4{2..3}}}b", "a{1}b a{42}b a{43}b"},
		{"1{b..e}2{a..c}3", "1b2a3 1b2b3 1b2c3 1c2a3 1c2b3 1c2c3 1d2a3 1d2b3 1d2c3 1e2a3 1e2b3 1e2c3"},
		{"as{12,32}{a..c}{2}", "as12a{2} as12b{2} as12c{2} as32a{2} as32b{2} as32c{2}"},
	}

	for _, tt := range tests {
		result := strings.Join(expandSingle(tt.in), space)
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}

func TestGetPair(t *testing.T) {
	type data struct {
		in    string
		start int
		stop  int
	}
	tests := []data{
		{"{x{12,{}}xxxxx", 2, 8},
		{"x{12,{}}{{,13}", 1, 7},
		{"{x{12,{}}{{,13}", 2, 8},
		{"{{x{12,{}}{{,13}", 3, 9},
		{"}some{", -1, -1},
		{"}some", -1, -1},
		{"some{", -1, -1},
	}

	for _, tt := range tests {
		start, stop := getPair(tt.in)
		assert.Equal(t, tt.start, start, "start of %s", tt.in)
		assert.Equal(t, tt.stop, stop, "stop of %s", tt.in)
	}
}
