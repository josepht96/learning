package main

import "fmt"

func PrintString(s string) string {
	fmt.Println(s)
	return s
}

func main() {
	s := PrintString("hello")
	fmt.Println(s)
}
