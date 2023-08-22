package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	type data struct {
		in  []byte
		out []byte
	}

	tests := []data{
		{[]byte("a{b c}d"), []byte("a{b c}d")},
		{[]byte("a{b,c}d"), []byte("abd acd")},
		{[]byte("a{b,c}d a{b c}d"), []byte("abd acd a{b c}d")},
		{[]byte("a{{b,c}d a{b c}d"), []byte("a{bd a{cd a{b c}d")},
		{[]byte("a{{b,c}}d a{b c}d"), []byte("a{b}d a{c}d a{b c}d")},
	}

	for _, tt := range tests {
		result := bytes.Join(Expand(tt.in), space)
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}

func TestExpandSingle(t *testing.T) {
	type data struct {
		in  []byte
		out []byte
	}

	tests := []data{
		{[]byte("{2..3}"), []byte("2 3")},
		{[]byte("a{{2..3}}b"), []byte("a{2}b a{3}b")},
		{[]byte("a{{{2..3}}}b"), []byte("a{{2}}b a{{3}}b")},
		{[]byte("a{{2}}b"), []byte("a{{2}}b")},
		{[]byte("a{{{2}}}b"), []byte("a{{{2}}}b")},
		{[]byte("a{{1,4{2..3}}}b"), []byte("a{1}b a{42}b a{43}b")},
		{[]byte("1{b..e}2{a..c}3"), []byte("1b2a3 1b2b3 1b2c3 1c2a3 1c2b3 1c2c3 1d2a3 1d2b3 1d2c3 1e2a3 1e2b3 1e2c3")},
		{[]byte("as{12,32}{a..c}{2}"), []byte("as12a{2} as12b{2} as12c{2} as32a{2} as32b{2} as32c{2}")},
	}

	for _, tt := range tests {
		result := bytes.Join(expandSingle(tt.in), space)
		assert.Equal(t, tt.out, result, "input %q", tt.in)
	}
}

func TestGetPair(t *testing.T) {
	type data struct {
		in    []byte
		start int
		stop  int
	}
	tests := []data{
		{[]byte("{x{12,{}}xxxxx"), 2, 8},
		{[]byte("x{12,{}}{{,13}"), 1, 7},
		{[]byte("{x{12,{}}{{,13}"), 2, 8},
		{[]byte("{{x{12,{}}{{,13}"), 3, 9},
		{[]byte("}some{"), -1, -1},
		{[]byte("}some"), -1, -1},
		{[]byte("some{"), -1, -1},
	}

	for _, tt := range tests {
		start, stop := getPair(tt.in)
		assert.Equal(t, tt.start, start, "start of %s", tt.in)
		assert.Equal(t, tt.stop, stop, "stop of %s", tt.in)
	}
}
