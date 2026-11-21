// First and last element of target
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(search(arr, target))
}

func search(arr []int, target int) []int {
	first := first(arr, target)
	last := last(arr, target)

	if first == -1 && last == -1 {
		return []int{-1, -1}
	}

	return []int{first, last}
} // time O(1), space O(1)

func first(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			right = mid - 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
} // time O(log n), space O(1)

func last(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
} // time O(log n), space O(1)

// total time := O(log n) + O(log n) = O(log n)
// space := O(1)
