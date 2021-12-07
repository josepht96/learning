package main

import "fmt"

type queue struct {
	queue []int
}

func (q queue) enqueue(num int) queue {
	fmt.Printf("Adding new number: %d\n", num)
	q.queue = append(q.queue, num)
	return q
}
func (q queue) dequeue() (int, queue) {
	n := q.queue[0]
	q.queue = q.queue[1:]
	return n, q
}
func queueStart() {
	q := queue{}
	fmt.Println(q.queue)
	q = q.enqueue(5)
	q = q.enqueue(4)
	q = q.enqueue(3)
	q = q.enqueue(2)
	q = q.enqueue(1)
	q = q.enqueue(56)
	fmt.Println(q.queue)
	n, q := q.dequeue()
	fmt.Println(n)
	fmt.Println(q.queue)
	n, q = q.dequeue()
	fmt.Println(n)
	fmt.Println(q.queue)
}
