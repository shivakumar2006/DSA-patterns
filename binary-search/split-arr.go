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
		left = max(left, num)
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

// first take left as the largest element in array and right as the total sum of the array
// now find mid
// then check if we can split the array in k subarrays with capacity mid
// if yes then right = mid
// else left = mid + 1

// in canSplit we check if we can split the array in k subarrays with capacity mid
// if yes then return true
// else return false
// time complexity of canSplit is O(n)
// space complexity of canSplit is O(1)
