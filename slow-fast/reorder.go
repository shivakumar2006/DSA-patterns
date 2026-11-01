// Reorder list
// https://leetcode.com/problems/reorder-list/?utm_source=chatgpt.com

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

func (l *LinkedList) middleNode() *Node {
	slow := l.head
	fast := l.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func (l *LinkedList) Reorder() {
	if l.head == nil || l.head.Next == nil {
		return
	}

	mid := l.middleNode()

	// split the list
	secondHead := mid.Next
	mid.Next = nil

	// reverse the second half
	var prev *Node
	current := secondHead

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	secondHead = prev

	// merge both halves
	firstHead := l.head

	for secondHead != nil {
		temp1 := firstHead.Next
		temp2 := secondHead.Next

		firstHead.Next = secondHead
		secondHead.Next = temp1

		firstHead = temp1
		secondHead = temp2
	}

	// update tail
	l.tail = firstHead

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

	fmt.Print("Original list : ")
	ll.Display()

	ll.Reorder()

	fmt.Print("reordered list : ")
	ll.Display()
}

// O(n) time
// O(1) space
