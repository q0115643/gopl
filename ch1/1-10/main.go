package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	run(1)
	run(2)
}

func run(i int) {
	var out *os.File
	var err error
	if out, err = os.Create("1-10-trial" + strconv.Itoa(i) + ".txt"); err != nil {
		fmt.Printf("while creating file: %v", err)
		return
	}
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		io.WriteString(out, <-ch)
	}
	io.WriteString(
		out,
		fmt.Sprintf("\n%.2fs elapsed\n", time.Since(start).Seconds()),
	)
	if err = out.Close(); err != nil {
		fmt.Printf("while closing file: %v", err)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
