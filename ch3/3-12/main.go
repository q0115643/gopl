package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input 2 strings")
		os.Exit(1)
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	if areAnagrams(s1, s2) {
		fmt.Println("They are anagrams")
	} else {
		fmt.Println("They are not anagrams")
	}
}

func areAnagrams(s1, s2 string) bool {
	s1Freq := make(map[rune]int)
	s2Freq := make(map[rune]int)
	for _, c := range s1 {
		s1Freq[c]++
	}
	for _, c := range s2 {
		s2Freq[c]++
	}
	for k, v := range s1Freq {
		if s2Freq[k] != v {
			return false
		}
	}
	for k, v := range s2Freq {
		if s1Freq[k] != v {
			return false
		}
	}
	return true
}
