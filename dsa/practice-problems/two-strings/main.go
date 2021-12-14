package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	s1 := flag.String("s1", "hello", "a string")
	s2 := flag.String("s2", "world", "a string")
	fmt.Printf("Outcome: %t\n", compareStrings(*s1, *s2))
}

func compareStrings(s1 string, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		fmt.Printf("string: %s\n", s1[:i+1])
		if strings.Contains(s2, s1[:i+1]) {
			fmt.Printf("%s is found in %s\n", s1[:i+1], s2)
			return true
		}
	}
	if len(s1) > 0 {
		if compareStrings(s1[1:], s2) {
			return true
		}
	}
	return false
}

//check every sequence of s1
