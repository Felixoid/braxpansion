# braxpansion
library for GO shell-like brace expansion

## Usage
It contains the only public functions `ExpandString(string) []string` and `ExpandBytes([]byte) [][]byte`, that expands shell-like expressions `1{b..e}2{a..c}3` to `["1b2a3" "1b2b3" "1b2c3" "1c2a3" "1c2b3" "1c2c3" "1d2a3" "1d2b3" "1d2c3" "1e2a3" "1e2b3" '1e2c3']`.

To see how it works and play around run:

```
go install go install github.com/Felixoid/braxpansion/braxpansion@latest
braxpansion '123 {} {1..012}{a..e}{first,last}'
```

## Why?
I couldn't find any descend library providing such functional with simply usage. Here are some benchmark results for [Braces](https://pkg.go.dev/mvdan.cc/sh@v2.6.4+incompatible/expand#Braces) and [gobrex](https://github.com/kujtimiihoxha/go-brace-expansion):

***Note***: Braces doesn't support `{z..A}` runes expansion.

```
$ { cd bench && go test -benchtime=2s -bench=. -benchmem . ; }
goos: linux
goarch: amd64
pkg: github.com/Felixoid/braxpansion/bench
cpu: AMD Ryzen 7 5700U with Radeon Graphics
BenchmarkExpand/Tiny-DummyWarmUp-16             34223894                72.26 ns/op           40 B/op          2 allocs/op
BenchmarkExpand/Tiny-ExpandString-16              885927              2637 ns/op            1776 B/op         51 allocs/op
BenchmarkExpand/Tiny-ExpandBytes-16               738478              3137 ns/op            2800 B/op         51 allocs/op
BenchmarkExpand/Tiny-ExpBytesAsStr-16             755206              3170 ns/op            2176 B/op         65 allocs/op
BenchmarkExpand/Tiny-Gobrex-16                     39058             61481 ns/op           51026 B/op        662 allocs/op
BenchmarkExpand/Tiny-ShExpand-16                  165868             14589 ns/op           16320 B/op        234 allocs/op
BenchmarkExpand/Small-DummyWarmUp-16            31715910                74.75 ns/op           40 B/op          2 allocs/op
BenchmarkExpand/Small-ExpandString-16             417553              5349 ns/op            4562 B/op         88 allocs/op
BenchmarkExpand/Small-ExpandBytes-16              369638              6303 ns/op            6540 B/op         90 allocs/op
BenchmarkExpand/Small-ExpBytesAsStr-16            362995              7056 ns/op            6147 B/op        130 allocs/op
BenchmarkExpand/Small-Gobrex-16                    25588             93668 ns/op           78098 B/op       1091 allocs/op
BenchmarkExpand/Small-ShExpand-16                 225572             10956 ns/op           14616 B/op        131 allocs/op
BenchmarkExpand/Big-DummyWarmUp-16              30555032                79.11 ns/op           40 B/op          2 allocs/op
BenchmarkExpand/Big-ExpandString-16               121489             19737 ns/op           15529 B/op        281 allocs/op
BenchmarkExpand/Big-ExpandBytes-16                 90355             25618 ns/op           27209 B/op        319 allocs/op
BenchmarkExpand/Big-ExpBytesAsStr-16               88429             26695 ns/op           23494 B/op        474 allocs/op
BenchmarkExpand/Big-Gobrex-16                      15571            152843 ns/op          111576 B/op       2110 allocs/op
BenchmarkExpand/Big-ShExpand-16                    18926            129354 ns/op           99260 B/op       2334 allocs/op
BenchmarkExpand/Huge-DummyWarmUp-16             26923173                88.93 ns/op           40 B/op          2 allocs/op
BenchmarkExpand/Huge-ExpandString-16               25449             95124 ns/op          115424 B/op       1059 allocs/op
BenchmarkExpand/Huge-ExpandBytes-16                18105            135030 ns/op          132641 B/op       1075 allocs/op
BenchmarkExpand/Huge-ExpBytesAsStr-16              17445            138515 ns/op          172373 B/op       2069 allocs/op
BenchmarkExpand/Huge-Gobrex-16                      4662            512728 ns/op          325247 B/op       7427 allocs/op
BenchmarkExpand/Huge-ShExpand-16                    2250           1038067 ns/op          850899 B/op      17043 allocs/op
PASS
ok      github.com/Felixoid/braxpansion/bench   67.339s
```
