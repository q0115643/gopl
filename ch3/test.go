package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p\n", &a)
	for i := range a {
		fmt.Printf("%p\n", &a[i])
	}
	fmt.Println("*****************")
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	fmt.Printf("%p\n", &a)
	for i := range a {
		fmt.Printf("%p\n", &a[i])
	}
	fmt.Println("*****************")
	b := "hello"
	fmt.Printf("%p\n", &b)
	b = "asdf"
	fmt.Printf("%p\n", &b)
	c := []byte(b)
	fmt.Printf("%p\n", &c)
	for i := 0; i < len(c); i++ {
		fmt.Printf("%p\n", &c[i])
	}
}
