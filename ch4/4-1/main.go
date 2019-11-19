package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	var cnt int = 0
	for i := 0; i < len(c1) && i < len(c2); i++ {
		if c1[i] != c2[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
