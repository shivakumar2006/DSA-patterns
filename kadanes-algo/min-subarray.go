// min subarray
// https://www.geeksforgeeks.org/problems/smallest-sum-contiguous-subarray/1

package main

import "fmt"

func main() {
	arr := []int{3, -4, 2, -3, -1, 7, -5}
	fmt.Println(minSubArray(arr))
}

func minSubArray(arr []int) int {
	currentSum := arr[0]
	minSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currentSum = min(arr[i], currentSum+arr[i])

		if currentSum < minSum {
			minSum = currentSum
		}
	}

	return minSum
}

// time O(n)
// space O(1)
