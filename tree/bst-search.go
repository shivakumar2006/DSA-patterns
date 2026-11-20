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

	for _, value := range values {
		root = insert(root, value)
	}

	fmt.Println("Binary search tree : ")
	display(root, 0)

	result := search(root, 40)

	if result != nil {
		fmt.Println("value found : ", result.Data)
	} else {
		fmt.Println("value not found")
	}

	fmt.Println("Search Result:")
	display(result, 0)
}
