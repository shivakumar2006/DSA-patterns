// binary search
// https://leetcode.com/problems/binary-search/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	result := bs(arr, target)
	fmt.Println(result)
}

func bs(arr []int, target int) int {
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
	return -1
}

// time O(log n)
// space O(1)
