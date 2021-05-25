# braxpansion
library for GO shell-like brace expansion

## Usage
It contains the only public function `Expand(string) []string`, that expands shell-like expressions `1{b..e}2{a..c}3` to `["1b2a3" "1b2b3" "1b2c3" "1c2a3" "1c2b3" "1c2c3" "1d2a3" "1d2b3" "1d2c3" "1e2a3" "1e2b3" '1e2c3']`.

## Why?
I couldn't find any descend library providing such functional with simply usage. Here are some benchmark results for [Braces](https://pkg.go.dev/mvdan.cc/sh@v2.6.4+incompatible/expand#Braces) and [gobrex](https://github.com/kujtimiihoxha/go-brace-expansion):

```
go test -benchtime=10s -bench=. -benchmem ./bench
goos: linux
goarch: amd64
pkg: github.com/Felixoid/braxpansion/bench
cpu: AMD Ryzen 7 4800H with Radeon Graphics
BenchmarkGobrex-16                 	  19832	   556680 ns/o  184319 B/op	   3315 allocs/op
BenchmarkShExpand-16               	 252078	    52750 ns/o   29984 B/op	    334 allocs/op
BenchmarkExpandString-16           	 183957	    61439 ns/o   29737 B/op	    439 allocs/op
BenchmarkExpandBytes-16            	 182880	    71469 ns/o   44049 B/op	    442 allocs/op
BenchmarkExpandBytesAsString-16    	 189706	    63163 ns/o   29801 B/op	    441 allocs/op
PASS
ok  	github.com/Felixoid/braxpansion/bench	69.274s
```
