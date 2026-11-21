// search insert position
// https://leetcode.com/problems/search-insert-position/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(search(arr, target))
}

func search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// time O(log n)
// space O(1)
