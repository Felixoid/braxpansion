# braxpansion
library for GO shell-like brace expansion

## Usage
It contains the only public function `Expand(string) []string`, that expands shell-like expressions `1{b..e}2{a..c}3` to `["1b2a3" "1b2b3" "1b2c3" "1c2a3" "1c2b3" "1c2c3" "1d2a3" "1d2b3" "1d2c3" "1e2a3" "1e2b3" '1e2c3']`.

## Why?
I couldn't find any descend library providing such functional with simply usage. Here are some benchmark results for [Braces](https://pkg.go.dev/mvdan.cc/sh@v2.6.4+incompatible/expand#Braces) and [gobrex](https://github.com/kujtimiihoxha/go-brace-expansion):

```
go test -benchtime=10s -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/Felixoid/braxpansion
cpu: AMD Ryzen 7 4800H with Radeon Graphics
BenchmarkExpand-16      	 192246	    68934 ns/op   20804 B/op	    613 allocs/op
BenchmarkGobrex-16      	  20702	   593651 ns/op  184433 B/op	   3315 allocs/op
BenchmarkShExpand-16    	 200557	    60172 ns/op   29984 B/op	    334 allocs/op
PASS
ok  	github.com/Felixoid/braxpansion	44.708s
```

## Plans
See if replacement of strings by []byte in internals would help to reduce allocations.
