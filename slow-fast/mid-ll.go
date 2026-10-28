// middle of the linked list
// https://leetcode.com/problems/middle-of-the-linked-list/description/

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

func (l *LinkedList) midNode() *Node {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
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

	mid := ll.midNode()

	fmt.Println("mid node of the list : ", mid.data)

	ll.Display()
}
