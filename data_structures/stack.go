package main

import "fmt"

type stack []int

func (s *stack) push(n int) stack {
	fmt.Printf("Adding top: %d\n", n)
	return append(*s, n)
}

func (s *stack) pop() stack {
	length := len(*s)
	fmt.Printf("Removing top: %d\n", (*s)[length-1])
	if length != 0 {
		return (*s)[:(length - 1)]
	}

	return *s
}

func stackStart() {
	s := stack{}
	s = s.push(1)
	s = s.push(2)
	s = s.pop()
}
