// search in rotated sorted array
// https://leetcode.com/problems/search-in-rotated-sorted-array/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(rotatedSorted(nums, target))
}

func rotatedSorted(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[left] <= arr[mid] {
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
} // time O(log n), space O(1)

// left = 0, right := len(arr) - 1
// find mid
// divide array into two parts
// check in first part from arr[left] to arr[mid]
// check target is greater or equals to start and less than mid
// right = mid - 1
// else {left = mid + 1}

// check right side of the sorted array from mid to end
// check target is greater than mid and less or equals to end
// left = mid + 1
// else {right = mid - 1}
