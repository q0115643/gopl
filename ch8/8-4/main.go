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

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	wg := sync.WaitGroup{}
	for input.Scan() {
		wg.Add(1)
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
