package main

import "fmt"

type dnode struct {
	name string
	prev *dnode
	next *dnode
}

type dlist struct {
	name string
	head *dnode
	tail *dnode
}

func (dl *dlist) newNode(name string) *dnode {
	newNode := &dnode{
		name: name,
		prev: nil,
		next: nil,
	}
	return newNode
}
func (dl *dlist) addNode(name string) {
	newNode := dl.newNode(name)
	if dl.head == nil {
		dl.head = newNode
	} else {
		currentNode := dl.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
		newNode.prev = currentNode
		newNode.next = nil
	}
}
func (dl *dlist) printList() {
	currentNode := dl.head
	for currentNode != nil {
		fmt.Printf("node: %s\n", currentNode.name)
		currentNode = currentNode.next
	}
}
func doubleLinkList() {
	dl := dlist{
		name: "dlist",
		head: nil,
		tail: nil,
	}
	dl.addNode("A")
	dl.addNode("B")
	dl.addNode("C")
	dl.addNode("D")
	dl.addNode("E")
	dl.addNode("F")
	dl.addNode("G")
	dl.addNode("H")
	dl.printList()
}
