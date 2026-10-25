// 2461. Maximum Sum of Distinct Subarrays With Length K
// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{1, 5, 4, 2, 9, 9, 9}
	k := 3
	fmt.Println(maximumSubarraySum(arr, k))
}

func maximumSubarraySum(arr []int, k int) int {
	windowSum := 0

	maxSum := 0
	freq := make(map[int]int)

	for i := 0; i < k; i++ {
		windowSum += arr[i]
		freq[arr[i]]++
	}

	if len(freq) == k {
		maxSum = windowSum
	}

	for i := k; i < len(arr); i++ {
		old := arr[i-k]
		freq[old]--
		if freq[old] == 0 {
			delete(freq, old)
		}

		new := arr[i]
		freq[new]++
		windowSum += new - old

		if len(freq) == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}
		}
	}

	return maxSum
}

// time O(n)
// space O(k) cause freq map has at most k distinct elements
