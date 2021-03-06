// Package braxpansion provides shell-like braces expansion.
//
// Examples:
// list of coma-separated arguments `v{aa,bb,cc}e` => `vaae vbbe vcce`
// single rune range `{a..f}` => `a b c d e f`
// numbers range `1{9..11}` => `19 110 111`
// numbers range with leading zeros `2{9..011}` => `2009 2010 2011` OR `2{09..11}` => `209 210 211`
// numbers range with increment
package braxpansion

import (
	"github.com/Felixoid/braxpansion/bytes"
	"github.com/Felixoid/braxpansion/strings"
)

// ExpandString takes a string contains a shell brace expansion expression and returns a slice of strings after
// they are expanded. As in shells, each word is processed separately, so `12{1,2,3,4}as ds{1..3}22` produces `121as 122as 123as 124as ds122 ds222 ds322`
func ExpandString(in string) []string {
	return strings.Expand(in)
}

// ExpandBytes takes a []byte contains a shell expansion expression and returns a slice of []byte after
// they are expanded. As in shells, each word is processed separately, so `12{1,2,3,4}as ds{1..3}22` produces `121as 122as 123as 124as ds122 ds222 ds322`.
// According to benchmark, it's better to convert []byte to string and use ExpandString.
func ExpandBytes(in []byte) [][]byte {
	return bytes.Expand(in)
}
