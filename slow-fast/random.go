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

func (l *LinkedList) hasCycle() bool {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}

	return false
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

func (l *LinkedList) detectCycle() *Node {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			slow = l.head

			for slow != fast {
				slow = slow.next
				fast = fast.next
			}

			return slow
		}
	}
	return nil
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

	cycleStart := ll.head.next
	ll.tail.next = cycleStart

	cycle := ll.hasCycle()
	length := ll.cycleLength()
	node := ll.detectCycle()

	fmt.Println("has cycle : ", cycle)

	fmt.Println("length of the cycle : ", length)

	if node != nil {
		fmt.Printf("cycle starts at node : %d\n", node.data)
	} else {
		fmt.Println("no cycle found")
	}

	ll.Display(9)
}
