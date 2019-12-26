package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type host struct {
	name    string
	address string
	port    string
}

func (h *host) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s=%s\n", h.name, s.Text())
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST ...")
		os.Exit(1)
	}
	hosts := make([]*host, 0)
	for _, str := range os.Args[1:] {
		split1 := strings.Split(str, "=")
		if len(split1) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", str)
			os.Exit(1)
		}
		split2 := strings.Split(split1[1], ":")
		if len(split2) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", str)
			os.Exit(1)
		}
		hosts = append(hosts, &host{split1[0], split2[0], split2[1]})
	}
	for _, host := range hosts {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host.address, host.port))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go host.watch(os.Stdout, conn)
	}
	for {
		time.Sleep(time.Minute)
	}
}
