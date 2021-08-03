package main

import "fmt"

type node struct {
	name string
	next *node
}
type list struct {
	name string
	head *node
	tail *node
}

func (l *list) createNode(name string) *node {
	newNode := &node{
		name: name,
		next: nil,
	}
	return newNode
}

func (l *list) addNode(name string) {
	newNode := l.createNode(name)
	if l.head == nil {
		l.head = newNode
		fmt.Printf("Head: %s\n", l.head.name)
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			fmt.Printf("Current node: %s\n", currentNode.next.name)
			currentNode = currentNode.next
		}
		fmt.Printf("Adding node %s, pointed to by %s\n", newNode.name, currentNode.name)
		l.tail = newNode
		currentNode.next = newNode
	}
}
func (l *list) addToTail(name string) {
	fmt.Printf("Current tail: %s\n", l.tail.name)
	newNode := l.createNode(name)
	l.tail.next = newNode
	l.tail = newNode

}
func (l *list) changeHead(name string) {
	newNode := l.createNode(name)
	newNode.next = l.head
	l.head = newNode
	fmt.Printf("Head changed to: %s\n", l.head.name)
}
func (l *list) removeNode(name string) {
	if l.head.name == name {
		l.head = l.head.next
		fmt.Printf("New head: %s\n", l.head.name)
	} else {
		prevNode := l.head
		currentNode := l.head.next
		for currentNode.name != name {
			prevNode = currentNode
			currentNode = currentNode.next
		}
		fmt.Printf("deleting node %s\n", currentNode.name)
		prevNode.next = currentNode.next
		currentNode.next = nil
	}
}

func (l *list) printList() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("node: %s\n", currentNode.name)
		currentNode = currentNode.next
	}
}
func linkList() {
	l := &list{
		name: "somelist",
		head: nil,
		tail: nil,
	}
	l.addNode("A")
	l.addNode("B")
	l.addNode("C")
	l.addNode("D")
	l.addNode("E")
	l.addNode("F")
	l.addNode("G")
	l.addNode("H")
	l.printList()
	l.removeNode("A")
	l.printList()
	l.addToTail("I")
	l.printList()
	l.changeHead("Z")
	l.printList()
}
