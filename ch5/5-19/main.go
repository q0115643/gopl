package main

import "fmt"

func recoverTest() (rst string) {
	defer func() {
		if p := recover(); p != nil {
			rst = fmt.Sprintf("PANIC: %s", p)
		}
	}()
	panic("T_T")
}

func main() {
	fmt.Println(recoverTest())
}
