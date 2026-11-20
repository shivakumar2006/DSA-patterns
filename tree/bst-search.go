// Search in a BST
// https://leetcode.com/problems/search-in-a-binary-search-tree/?utm_source=chatgpt.com

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Data: val}
	}

	if val < root.Data {
		root.Left = insert(root.Left, val)
	} else if val > root.Data {
		root.Right = insert(root.Right, val)
	}

	return root
}

func search(root *Node, val int) *Node {
	if root == nil {
		return nil
	}

	if root.Data == val {
		return root
	}

	if val < root.Data {
		return search(root.Left, val)
	}

	return search(root.Right, val)
}

func display(root *Node, level int) {
	if root == nil {
		return
	}

	display(root.Right, level+1)

	for i := 0; i < level; i++ {
		fmt.Print("   ")
	}

	fmt.Println(root.Data)

	display(root.Left, level+1)
}

func main() {
	values := []int{50, 30, 70, 20, 40, 60, 80}

	var root *Node

	// Create BST
	for _, value := range values {
		root = insert(root, value)
	}

	fmt.Println("Original Binary Search Tree:")
	display(root, 0)

	// Search value
	target := 30
	result := search(root, target)

	if result != nil {
		fmt.Println("\nValue found:", target)
		fmt.Println("Subtree:")
		display(result, 0)
	} else {
		fmt.Println("\nValue not found:", target)
	}
}
