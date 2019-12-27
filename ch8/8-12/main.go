package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	channel chan<- string
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.channel <- "---All Clients---"
			for cli2 := range clients {
				cli.channel <- cli2.name
			}
			cli.channel <- "-----------------"

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	name := conn.RemoteAddr().String()
	cli := client{ch, name}
	go clientWriter(conn, ch)

	cli.channel <- "You are " + cli.name
	messages <- cli.name + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- cli.name + ": " + input.Text()
	}

	leaving <- cli
	messages <- cli.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
