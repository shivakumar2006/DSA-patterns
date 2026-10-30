// reverse linked list 2
// https://github.com/shivakumar2006/DSA-go/blob/main/LinkedList/Interview/reversell2.go

package main

import "fmt"

type Node struct {
	data int
	Next *Node
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

	l.tail.Next = newNode
	l.tail = newNode
}

func (l *LinkedList) ReverseBetween(start, end int) *Node {
	if start == end {
		return l.head
	}

	var before *Node
	current := l.head

	// reach the start position
	for i := 0; i < start; i++ {
		before = current
		current = current.Next
	}

	// store the first node of the portion
	first := current

	var prev *Node

	for i := start; i <= end; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// reconnect the reverse portion
	if before != nil {
		before.Next = prev
	} else {
		l.head = prev
	}

	first.Next = current

	// update tail if reversal reaches the end
	if current == nil {
		l.tail = first
	}

	return l.head
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.Next
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

	fmt.Println("Original list : ")
	ll.Display()

	ll.head = ll.ReverseBetween(1, 3)

	fmt.Println("Reversed list : ")
	ll.Display()
}
