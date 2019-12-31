package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := 1
	out := make(chan int)
	first := out
	prev := time.Now()
	defer func() {
		cur := time.Now()
		fmt.Println(cur.Sub(prev))
		fmt.Println(cnt)
	}()
	for ; cnt < 1000000000000000; cnt++ {
		in := out
		out = make(chan int)
		go func(in, out chan int) {
			for v := range in {
				out <- v
			}
		}(in, out)
	}
	first <- 1
	<-out
}
