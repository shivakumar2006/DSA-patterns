// Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/description/?utm_source=chatgpt.com

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("end")
}

func main() {
	ll := &LinkedList{}

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)

	ll.Display()

	ll.Reverse()
	ll.Display()
}
