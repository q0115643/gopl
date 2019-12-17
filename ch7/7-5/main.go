package main

import (
	"fmt"
	"io"
	"strings"
)

type IOLimitReader struct {
	r        io.Reader
	n, limit int
}

func (r *IOLimitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n > r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &IOLimitReader{r: r, limit: limit}
}

func main() {
	r := LimitReader(strings.NewReader("12345678901234567890123456789012345678901234567890"), 40)
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	buffer = buffer[:n]
	fmt.Println(n, err, string(buffer))
}
