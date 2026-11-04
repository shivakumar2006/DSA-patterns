// max subarray
// https://leetcode.com/problems/maximum-subarray/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(arr))
}

func maxSubArray(arr []int) int {
	currentSum := arr[0]
	maxSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currentSum = max(arr[i], currentSum+arr[i])

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// time O(n)
// space O(1)
