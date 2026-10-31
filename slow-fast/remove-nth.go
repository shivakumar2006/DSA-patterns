// remove Nth node from end of list
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/?utm_source=chatgpt.com

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

func (l *LinkedList) RemoveNthFromEnd(n int) {
	slow := l.head
	fast := l.head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		l.head = l.head.Next

		if l.head == nil {
			l.tail = nil
		}

		return
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	if slow.Next == nil {
		l.tail = slow
	}
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
	n := 2

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)

	ll.Display()

	ll.RemoveNthFromEnd(n)

	ll.Display()
}
