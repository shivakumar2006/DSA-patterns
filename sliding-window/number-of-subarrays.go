// 1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k := 3
	threshold := 4
	fmt.Println(numOfSubarrays(arr, k, threshold))
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	windowSum := 0
	count := 0
	target := k * threshold

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	if windowSum >= target {
		count++
	}

	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum >= target {
			count++
		}
	}

	return count
}

// time O(n)
// space O(1)
