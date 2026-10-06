// minimum size subarray sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/

package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArray(arr, target))
}

func minSubArray(arr []int, target int) int {
	left := 0
	sum := 0
	minLen := len(arr) + 1

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		for sum >= target {
			length := right - left

			if length < minLen {
				minLen = length
			}

			sum -= arr[left]
			left++
		}
	}

	if minLen == len(arr)+1 {
		return 0
	}

	return minLen
}

// func minSubArray(arr []int, target int) int {
// 	left := 0
// 	sum := 0
// 	minLen := len(arr) + 1

// 	for right := 0; right < len(arr); right++ {
// 		sum += arr[right]

// 		for sum >= target {
// 			length := right - left + 1

// 			if length < minLen {
// 				minLen = length
// 			}

// 			sum -= arr[left]
// 			left++
// 		}
// 	}

// 	if minLen == len(arr)+1 {
// 		return 0
// 	}

// 	return minLen
// }

// time O(n)
// space O(1)

// func minSubArray(arr []int, target int) int {
// 	left := 0
// 	sum := 0
// 	minLen := len(arr) + 1

// 	for right := 0; right < len(arr); right++ {
// 		sum += arr[right]

// 		for sum >= target {
// 			length := right - left + 1

// 			if length < minLen {
// 				minLen = length
// 			}

// 			sum -= arr[left]
// 			left++
// 		}
// 	}

// 	if minLen == len(arr)+1 {
// 		return 0
// 	}

// 	return minLen
// }
