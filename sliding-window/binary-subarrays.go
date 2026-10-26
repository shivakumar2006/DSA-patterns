// binary subarray with sum
// https://leetcode.com/problems/binary-subarrays-with-sum/submissions/2182127450/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{1, 0, 1, 0, 1}
	goal := 2
	fmt.Println(numSubarraysWithSum(arr, goal))
}

func numSubarraysWithSum(arr []int, goal int) int {
	return atMost(arr, goal) - atMost(arr, goal-1)
}

func atMost(arr []int, goal int) int {
	if goal < 0 {
		return 0
	}

	left := 0
	sum := 0
	count := 0

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		for sum > goal {
			sum -= arr[left]
			left++
		}

		length := right - left + 1
		count += length
	}

	return count
}
