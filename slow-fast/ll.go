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

func (l *LinkedList) Insert(val int) {
	newNode := &Node{data: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	l.tail = newNode
}

func (l *LinkedList) InsertLast(val int) {
	newNode := &Node{data: val}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) InsertAt(val, index int) {
	newNode := &Node{data: val}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}

	temp := l.head
	for i := 1; i < index; i++ {
		temp = temp.next
	}

	if temp == nil {
		fmt.Println("index out of bounds")
		return
	}

	// insert new node
	newNode.next = temp.next
	temp.next = newNode

	// update tail if inserted at the end
	if newNode.next == nil {
		l.tail = newNode
	}
}

func (l *LinkedList) DeleteFirst() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	}
}

func (l *LinkedList) DeleteLast() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	temp := l.head
	for temp.next != l.tail {
		temp = temp.next
	}

	temp.next = nil
	l.tail = temp
}

func (l *LinkedList) DeleteAt(index int) {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	if index == 0 {
		l.head = l.head.next
		return
	}

	current := l.head
	for i := 0; current != nil && i < index-1; i++ {
		current = current.next
	}

	// if index if out of range
	if current == nil || current.next == nil {
		fmt.Println("index out of range")
	}

	current.next = current.next.next
	return
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
	ll := LinkedList{}
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)

	ll.InsertLast(40)
	ll.InsertLast(50)
	ll.InsertLast(60)

	ll.InsertAt(100, 3)

	ll.DeleteFirst()
	ll.DeleteLast()
	ll.DeleteAt(2)

	ll.Display()
}
