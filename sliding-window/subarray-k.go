// max sum subarray of size k
// https://www.geeksforgeeks.org/problems/max-sum-subarray-of-size-k5313/1

package main

import "fmt"

func main() {
	arr := []int{100, 200, 300, 400}
	k := 2
	fmt.Println(maxSum(arr, k))
}

func maxSum(arr []int, k int) int {
	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	// slide the window
	for i := k; i < len(arr); i++ {
		windowSum += arr[i]   // new element added
		windowSum -= arr[i-k] // old element remove

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// time O(n)
// space O(1)
