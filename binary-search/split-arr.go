// 410. Split Array Largest Sum
// https://leetcode.com/problems/split-array-largest-sum/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{7, 2, 5, 10, 8}
	k := 2
	fmt.Println(splitArr(arr, k))
}

func splitArr(arr []int, k int) int {
	left := 0
	right := 0

	for _, num := range arr {
		if num > left {
			left = num
		}
		right += num
	}

	for left < right {
		mid := left + (right-left)/2
		if canSplit(arr, mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
} // time O(log s) s is the sum of array and space O(1)
// Overall complexity O(n log s)

func canSplit(arr []int, capacity, k int) bool {
	subarrays := 1
	currentSum := 0

	for _, num := range arr {
		if currentSum+num > capacity {
			subarrays++
			currentSum = 0
		}
		currentSum += num
	}
	return subarrays <= k
} // time O(n), space O(1)
