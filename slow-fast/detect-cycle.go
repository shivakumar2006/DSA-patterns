// detect cycle in linked list
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

func (l *LinkedList) insert(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
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

func (l *LinkedList) display(limit int) {
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

	ll.insert(3)
	ll.insert(2)
	ll.insert(0)
	ll.insert(-4)

	// Create cycle: -4 -> 2
	cycleStart := ll.head.next
	ll.tail.next = cycleStart

	node := ll.detectCycle()

	if node != nil {
		fmt.Printf("Cycle starts at node %d\n", node.data)
	} else {
		fmt.Println("No cycle found")
	}

	ll.display(10)
}
