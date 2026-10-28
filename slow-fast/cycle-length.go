// https://leetcode.com/problems/linked-list-cycle-ii/description/?utm_source=chatgpt.com

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

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	l.tail = newNode
}

func (l *LinkedList) cycleLength() int {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			temp := slow.next
			length := 1

			for temp != slow {
				temp = temp.next
				length++
			}

			return length
		}
	}

	return 0
}

func (l *LinkedList) Display(limit int) {
	temp := l.head
	count := 0

	for temp != nil && count < limit {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
		count++
	}
	fmt.Println("...")
}

func main() {
	ll := &LinkedList{}

	ll.Insert(3)
	ll.Insert(2)
	ll.Insert(0)
	ll.Insert(-4)

	temp := ll.head
	var thirdNode *Node

	for temp.next != nil {
		if temp.data == 2 {
			thirdNode = temp
		}
		temp = temp.next
	}

	temp.next = thirdNode

	length := ll.cycleLength()

	fmt.Println("cycle length : ", length)

	ll.Display(7)
}
