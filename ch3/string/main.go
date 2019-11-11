package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	s := "hesadfasdfsallo"
	fmt.Println(len(s))
	fmt.Println(unsafe.Sizeof(s))
	d := s
	fmt.Println(s)
	fmt.Println(&s)
	ss := s[1:]
	fmt.Println(&ss)

	fmt.Println(d)
	fmt.Println(&d)
	s += "asasdfasdfasfdsafsafsadfasfasddf"
	d += "zxcv"
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Println(d)
	fmt.Println(&d)
	f := s
	fmt.Println(f)
	fmt.Println(&f)

	fmt.Println(len(s))
	fmt.Println(unsafe.Sizeof(s))

	fmt.Println("********************")
	s = "asdf"
	s += "asdfasdf"
	for i := range s {
		fmt.Printf("%c\n", s[i])
	}

	fmt.Println("********************")
	s = "김형록"
	fmt.Println(len(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println("********************")
	s = "김형록"
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d\t%c\n", i, s[i])
	}
	fmt.Println("********************")
	s = "김형록"
	fmt.Println(len(s))
	for i, c := range s {
		fmt.Printf("%d\t%c\n", i, c)
	}
	fmt.Println("********************")
	s = "asdf김형록"
	p := []rune(s)
	fmt.Println(len(p))
	for i, c := range p {
		fmt.Printf("%d\t%c\n", i, c)
	}
	fmt.Println("********************")
	s = "asdf김형록"
	fmt.Printf("%T\n", s)
	fmt.Println(&s)
	vb := s[1:]
	fmt.Printf("%T\n", vb)
	fmt.Println(&vb)
	fmt.Println("********************")
	k := [10]int64{1, 2222, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(&k)
	fmt.Println(&k[0])
	fmt.Println(&k[1])
	fmt.Println(&k[2])
	fmt.Println("********************")
	s = "asdf김형록"
	fmt.Println(&s)
	check(s)
	fmt.Printf("%p\n", &k[0])
	check2(k)
}

func check(s string) {
	fmt.Println(&s)
}

func check2(k [10]int64) {
	fmt.Printf("%p\n", &k[0])
}
