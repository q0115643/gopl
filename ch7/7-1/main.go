package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	cnt := 0
	for scanner.Scan() {
		cnt++
	}
	*c += WordCounter(cnt)
	return cnt, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	cnt := 1
	for _, b := range p {
		if b == '\n' {
			cnt++
		}
	}
	*c += LineCounter((cnt))
	return cnt, nil
}

func main() {
	input := "Hello World\nit's rokokss\nend"
	var wc WordCounter
	wc.Write([]byte(input))
	fmt.Println(wc)
	var lc LineCounter
	lc.Write([]byte(input))
	fmt.Println(lc)
}
