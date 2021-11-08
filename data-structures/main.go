package main

import "fmt"

//array +
//linked list +
//stack +
//queue
//maps
//graphs
//trees
//binary tree
//heaps
//self balancing tree

type test struct {
}

func main() {
	//pointer()
	linkList()
	//doubleLinkList()
	//stackStart()
	//queueStart()
}

func pointer() {
	num := 5
	fmt.Println(&num)
	var p1 *int = &num
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(&p1)
	var p2 *int = &num
	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println(&p2)
	*p2 = 6
	fmt.Println(num)
}
