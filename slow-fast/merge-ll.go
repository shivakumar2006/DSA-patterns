// merge two sorted list
// https://leetcode.com/problems/merge-two-sorted-lists/description/?utm_source=chatgpt.com

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

func (l *LinkedList) Merge(l1, l2 *LinkedList) *LinkedList {
	first := l1.head
	second := l2.head

	result := &LinkedList{}

	for first != nil && second != nil {
		if first.data < second.data {
			result.Insert(first.data)
			first = first.next
		} else {
			result.Insert(second.data)
			second = second.next
		}
	}

	for first != nil {
		result.Insert(first.data)
		first = first.next
	}

	for second != nil {
		result.Insert(second.data)
		second = second.next
	}

	return result
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
	l1 := &LinkedList{}
	l1.Insert(1)
	l1.Insert(2)
	l1.Insert(4)

	l2 := &LinkedList{}
	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	l1.Display()
	l2.Display()

	merged := &LinkedList{}
	merged = merged.Merge(l1, l2)

	merged.Display()
}
