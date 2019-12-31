package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ping := make(chan int)
	pong := make(chan int)
	var i int64
	start := time.Now()
	defer func() {
		fmt.Println(float64(i) / float64(time.Since(start)) * 1e9)
	}()
	go func() {
		ping <- 1
		for {
			i++
			ping <- <-pong
		}
	}()
	go func() {
		for {
			pong <- <-ping
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
