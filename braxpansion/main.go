package main

import (
	"fmt"
	"os"
	"strings"

	brStr "github.com/Felixoid/braxpansion/strings"
)

func main() {
	out := make([]string, 0, len(os.Args)-1)
	for _, a := range os.Args[1:] {
		out = append(out, brStr.Expand(a)...)
	}
	fmt.Println(strings.Join(out, " "))
}
