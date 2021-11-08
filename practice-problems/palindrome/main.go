package main

import (
	"flag"
	"fmt"
	"strings"
)

func SpaceFieldsJoin(str *string) {
	*str = strings.Join(strings.Fields(*str), "")
}

func main() {
	//could also check from punctuation and letter case
	s := flag.String("input string", "red rum sir is murder", "string to test")
	SpaceFieldsJoin(s)
	if isPalindrome(*s) {
		fmt.Printf("True\n")
	} else {
		fmt.Printf("False\n")
	}
}

func isPalindrome(s string) bool {
	fmt.Printf("S: %s\n", s)
	for i := 0; i < len(s); i++ {
		if s[i] != s[(len(s)-1)-i] {
			return false
		}
	}
	return true
}
