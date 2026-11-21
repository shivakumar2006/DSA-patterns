// Duplicate rotated sorted
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(arr, target))
}

func search(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		}

		if arr[left] == arr[mid] && arr[mid] == arr[right] {
			left++
			right--
			continue
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

	return false
}
