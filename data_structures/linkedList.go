package main

import "fmt"

type node struct {
	name string
	next *node
}
type list struct {
	name string
	head *node
}

func (l *list) createNode(name string) node {
	newNode := node{
		name: name,
		next: nil,
	}
	return newNode
}

func (l *list) addNode(name string) {
	fmt.Println("Here")
	newNode := l.createNode(name)
	if l.head == nil {
		l.head = &newNode
		fmt.Printf("Head: %s\n", l.head.name)
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			fmt.Printf("Current node: %s\n", currentNode.next.name)
			currentNode = currentNode.next
		}
		fmt.Printf("Adding node %s, pointed to by %s\n", newNode.name, currentNode.name)
		currentNode.next = &newNode
	}
}

func linkListStart() {
	l := &list{
		name: "somelist",
		head: nil,
	}
	l.addNode("A")
	l.addNode("B")
	l.addNode("C")
	l.addNode("D")
	l.addNode("E")
	l.addNode("F")
	l.addNode("G")
	l.addNode("H")
}
