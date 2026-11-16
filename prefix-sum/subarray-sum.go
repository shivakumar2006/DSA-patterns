// subarray sum equals k
// https://leetcode.com/problems/subarray-sum-equals-k/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(arr, k))
}

func subarraySum(arr []int, k int) int {
	prefixSum := 0
	count := 0

	freq := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		prefixSum += arr[i]

		if prefixSum == k {
			count++
		}

		if freq[prefixSum-k] > 0 {
			count += freq[prefixSum-k]
		}

		freq[prefixSum]++
	}

	return count
}

// time O(n)
// space O(n)
