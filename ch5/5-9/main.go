package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 0 {
		fmt.Printf("%s\n", expand(os.Args[1], toAbc))
	}
}

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$foo", f("foo"))
}

func toAbc(s string) string {
	return "abc"
}
