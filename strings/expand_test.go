package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandSingle(t *testing.T) {
	type data struct {
		in  string
		out string
	}

	tests := []data{
		{in: "1{b..e}2{a..c}3", out: "1b2a3 1b2b3 1b2c3 1c2a3 1c2b3 1c2c3 1d2a3 1d2b3 1d2c3 1e2a3 1e2b3 1e2c3"},
		{in: "as{12,32}{a..c}{2}", out: "as12a{2} as12b{2} as12c{2} as32a{2} as32b{2} as32c{2}"},
	}

	for _, tt := range tests {
		result, err := expandSingle(tt.in)
		assert.NoError(t, err)
		assert.Equal(t, tt.out, strings.Join(result, space), "input %q", tt.in)
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
	}

	for _, tt := range tests {
		start, stop := getPair(tt.in)
		assert.Equal(t, tt.start, start, "start of %s", tt.in)
		assert.Equal(t, tt.stop, stop, "stop of %s", tt.in)
	}
}
