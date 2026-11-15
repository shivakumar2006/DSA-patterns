// max absolute sum of any subarray
// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/description/

package main

import "fmt"

func main() {
	arr := []int{1, -3, 2, 3, -4}
	fmt.Println(maxAbsSum(arr))
}

func maxAbsSum(arr []int) int {
	maxSum := 0
	minSum := 0

	maxCurrent := 0
	minCurrent := 0

	for i := 0; i < len(arr); i++ {
		maxCurrent = max(arr[i], maxCurrent+arr[i])
		minCurrent = min(arr[i], minCurrent+arr[i])

		maxSum = max(maxSum, maxCurrent)
		minSum = min(minSum, minCurrent)
	}

	return max(maxSum, abs(minSum))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// time O(n)
// space O(1)
