package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func one() {
	t := time.Now()
	var s, sep string
	for i, arg := range os.Args[:] {
		s += sep + strconv.Itoa(i) + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
	fmt.Println(time.Since(t) * 100000000000)
}

func two() {
	t := time.Now()
	var s, sep string
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(t) * 100000000000)
}

func three() {
	t := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(time.Since(t) * 100000000000)

}

func main() {
	one()
	two()
	three()
}
