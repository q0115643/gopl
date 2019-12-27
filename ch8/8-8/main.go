package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func countDown(c net.Conn, typedIn <-chan bool) {
	tick := time.NewTicker(1 * time.Second)
	cnt := 0
	for {
		select {
		case <-tick.C:
			cnt++
			if cnt == 5 {
				fmt.Fprintln(c, "kick")
				tick.Stop()
				c.Close()
				return
			}
		case <-typedIn:
			cnt = 0
		}
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	typedIn := make(chan bool)
	wg := sync.WaitGroup{}
	go countDown(c, typedIn)
	for input.Scan() {
		wg.Add(1)
		typedIn <- true
		go func(c net.Conn, shout string, delay time.Duration) {
			defer wg.Done()
			echo(c, shout, delay)
		}(c, input.Text(), 1*time.Second)
	}
	wg.Wait()
	if tcpconn, ok := c.(*net.TCPConn); ok {
		tcpconn.CloseWrite()
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
