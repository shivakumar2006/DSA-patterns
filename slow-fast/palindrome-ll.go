// palindrome linked list
// https://leetcode.com/problems/palindrome-linked-list/description/?utm_source=chatgpt.com

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

func (l *LinkedList) isPalindrome() bool {
	if l.head == nil || l.head.Next == nil {
		return true
	}

	slow := l.head
	fast := l.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	current := second
	var prev *Node

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	first := l.head
	second = prev

	for second != nil {
		if first.data != second.data {
			return false
		}

		first = first.Next
		second = second.Next
	}

	return true
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
	ll.Insert(2)
	ll.Insert(1)

	ll.Display()

	fmt.Println("is palindrome : ", ll.isPalindrome())
}

// package main

// import "fmt"

// type Node struct {
// 	data int
// 	Next *Node
// }

// type LinkedList struct {
// 	head *Node
// 	tail *Node
// }

// func (l *LinkedList) Insert(value int) {
// 	newNode := &Node{data: value}

// 	if l.head == nil {
// 		l.head = newNode
// 		l.tail = newNode
// 		return
// 	}

// 	l.tail.Next = newNode
// 	l.tail = newNode
// }

// func (l *LinkedList) isPalindrome() bool {
// 	if l.head == nil || l.head.Next == nil {
// 		return true
// 	}

// 	slow := l.head
// 	fast := l.head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	// reverse the second half
// 	second := slow.Next
// 	current := second
// 	var prev *Node

// 	for current != nil {
// 		next := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}

// 	// compare the first half with the reversed second half
// 	first := l.head
// 	second = prev

// 	for second != nil {
// 		if first.data != second.data {
// 			return false
// 		}

// 		first = first.Next
// 		second = second.Next
// 	}

// 	return true
// }

// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.Next
// 	}
// 	fmt.Println("end")
// }

// func main() {
// 	ll := &LinkedList{}

// 	ll.Insert(1)
// 	ll.Insert(2)
// 	ll.Insert(3)
// 	ll.Insert(2)
// 	ll.Insert(1)

// 	ll.Display()

// 	fmt.Println("is palindrome : ", ll.isPalindrome())
// }
