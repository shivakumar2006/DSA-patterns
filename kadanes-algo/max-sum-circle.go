// maximum sum circular subarray
// https://leetcode.com/problems/maximum-sum-circular-subarray/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{1, -2, 3, -2}
	fmt.Println(maxSubarraySumCircular(arr))
}

func maxSubarraySumCircular(arr []int) int {
	totalSum := 0

	maxSum := arr[0]
	minSum := arr[0]

	maxCurrent := 0
	minCurrent := 0

	for i := 0; i < len(arr); i++ {
		totalSum += arr[i]

		maxCurrent = max(arr[i], maxCurrent+arr[i])
		maxSum = max(maxSum, maxCurrent)

		minCurrent = min(arr[i], minCurrent+arr[i])
		minSum = min(minSum, minCurrent)
	}

	if maxSum < 0 {
		return maxSum
	}

	circularSum := totalSum - minSum

	return max(maxSum, circularSum)
}

// time O(n)
// space O(1)
