// rotate list
// https://leetcode.com/problems/rotate-list/description/?utm_source=chatgpt.com

package main

import (
	"fmt"
)

type Node struct {
	data int
	Next *Node
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

	l.tail.Next = newNode
	l.tail = newNode
}

func (l *LinkedList) Rotate(k int) {
	if l.head == nil || l.head.Next == nil || k == 0 {
		return
	}

	current := l.head
	length := 0

	for current != nil {
		length++
		current = current.Next
	}

	// reduce unnecessary rotation
	k = k % length

	if k == 0 {
		return
	}

	// find new tail
	newTail := length - k
	current = l.head

	for i := 1; i < newTail; i++ {
		current = current.Next
	}

	newHead := current.Next
	l.tail.Next = l.head
	current.Next = nil

	l.head = newHead
	l.tail = current
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
	k := 2

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)

	ll.Display()

	fmt.Print("after rotation : ")
	ll.Rotate(k)
	ll.Display()
}

// time O(n)
// space O(1)
