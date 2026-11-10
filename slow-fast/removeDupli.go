// Remove duplicates from sorted list
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/?utm_source=chatgpt.com

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

func (l *LinkedList) oddEven() *Node {
	odd := l.head
	even := l.head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
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

	ll.Display()

	fmt.Println("odd even list : ", ll.oddEven())
	ll.Display()
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

// func (l *LinkedList) Insert(val int) {
// 	newNode := &Node{data: val}

// 	if l.head == nil {
// 		l.head = newNode
// 		l.tail = newNode
// 		return
// 	}

// 	l.tail.Next = newNode
// 	l.tail = newNode
// }

// func (l *LinkedList) remove() *Node {
// 	current := l.head

// 	for current != nil && current.Next != nil {
// 		if current.data == current.Next.data {
// 			current.Next = current.Next.Next
// 		} else {
// 			current = current.Next
// 		}
// 	}

// 	return current
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
// 	ll.Insert(1)
// 	ll.Insert(2)
// 	ll.Insert(3)
// 	ll.Insert(3)
// 	ll.Insert(4)
// 	ll.Insert(4)
// 	ll.Insert(4)
// 	ll.Insert(5)

// 	ll.Display()

// 	ll.remove()
// 	ll.Display()
// }

// time O(n)
// space O(1)
