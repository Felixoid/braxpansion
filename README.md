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
go test -benchtime=2s -bench=. -benchmem ./bench
goos: linux
goarch: amd64
pkg: github.com/Felixoid/braxpansion/bench
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkExpand/Tiny-DummyWarmUp-16         	15473564	      153.3 ns/op	     40 B/op	      2 allocs/op
BenchmarkExpand/Tiny-ExpandString-16        	 395733	     6538 ns/op	   1776 B/op	     51 allocs/op
BenchmarkExpand/Tiny-ExpandBytes-16         	 346645	     7991 ns/op	   2800 B/op	     51 allocs/op
BenchmarkExpand/Tiny-ExpBytesAsStr-16       	 293335	     8277 ns/op	   2176 B/op	     65 allocs/op
BenchmarkExpand/Tiny-Gobrex-16              	  16880	   141999 ns/op	  50623 B/op	    662 allocs/op
BenchmarkExpand/Tiny-ShExpand-16            	  70768	    35712 ns/op	  16320 B/op	    234 allocs/op
BenchmarkExpand/Small-DummyWarmUp-16        	8287765	      310.1 ns/op	     80 B/op	      4 allocs/op
BenchmarkExpand/Small-ExpandString-16       	  42933	    55608 ns/op	  29737 B/op	    439 allocs/op
BenchmarkExpand/Small-ExpandBytes-16        	  34615	    71417 ns/op	  44049 B/op	    442 allocs/op
BenchmarkExpand/Small-ExpBytesAsStr-16      	  27488	    86896 ns/op	  42096 B/op	    745 allocs/op
BenchmarkExpand/Small-Gobrex-16             	   4880	   573179 ns/op	 184271 B/op	   3315 allocs/op
BenchmarkExpand/Small-ShExpand-16           	  36319	    63914 ns/op	  30936 B/op	    365 allocs/op
BenchmarkExpand/Big-DummyWarmUp-16          	5884639	      429.7 ns/op	    120 B/op	      6 allocs/op
BenchmarkExpand/Big-ExpandString-16         	  32799	    82344 ns/op	  34628 B/op	    553 allocs/op
BenchmarkExpand/Big-ExpandBytes-16          	  24505	   102857 ns/op	  53423 B/op	    596 allocs/op
BenchmarkExpand/Big-ExpBytesAsStr-16        	  20901	   115396 ns/op	  48388 B/op	    903 allocs/op
BenchmarkExpand/Big-Gobrex-16               	   3078	   838374 ns/op	 256363 B/op	   4733 allocs/op
BenchmarkExpand/Big-ShExpand-16             	   9024	   255849 ns/op	  91384 B/op	   1705 allocs/op
BenchmarkExpand/Huge-DummyWarmUp-16         	4016671	      589.7 ns/op	    160 B/op	      8 allocs/op
BenchmarkExpand/Huge-ExpandString-16        	   8079	   319763 ns/op	 150049 B/op	   1612 allocs/op
BenchmarkExpand/Huge-ExpandBytes-16         	   5944	   443943 ns/op	 186062 B/op	   1673 allocs/op
BenchmarkExpand/Huge-ExpBytesAsStr-16       	   5928	   449989 ns/op	 220752 B/op	   2972 allocs/op
BenchmarkExpand/Huge-Gobrex-16              	   1398	  2043009 ns/op	 545477 B/op	  12158 allocs/op
BenchmarkExpand/Huge-ShExpand-16            	    757	  3173720 ns/op	 926158 B/op	  18747 allocs/op
PASS
ok  	github.com/Felixoid/braxpansion/bench	70.588s
```
