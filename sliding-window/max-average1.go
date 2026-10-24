// maximum average subarray 1

package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(arr []int, k int) float64 {
	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}

	maxSum := currentSum

	for i := k; i < len(arr); i++ {
		currentSum += arr[i]
		currentSum -= arr[i-k]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return float64(maxSum) / float64(k)
}

// time O(n)
// space O(1)
