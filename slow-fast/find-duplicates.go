// find the duplicate number
// https://leetcode.com/problems/find-the-duplicate-number/description/

package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Println(find(arr))
}

func find(arr []int) int {
	slow := arr[0]
	fast := arr[0]

	for {
		slow = arr[slow]
		fast = arr[arr[fast]]

		if slow == fast {
			break
		}
	}

	slow = arr[0]

	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
	}

	return slow
}
