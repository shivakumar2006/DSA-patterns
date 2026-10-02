// triplets with smaller sum
// https://www.geeksforgeeks.org/problems/count-triplets-with-sum-smaller-than-x5549/1

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-2, 0, 1, 3}
	target := 2
	fmt.Println(smaller(arr, target))
}

func smaller(arr []int, target int) int {
	sort.Ints(arr)

	count := 0

	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum < target {
				count += right - left
				left++
			}
			if sum >= target {
				right--
			}
		}
	}
	return count
}

// func smaller(arr []int, target int) int {
// 	sort.Ints(arr)

// 	count := 0

// 	for i := 0; i < len(arr)-2; i++ {
// 		left := i + 1
// 		right := len(arr) - 1

// 		for left < right {
// 			sum := arr[i] + arr[left] + arr[right]

// 			if sum < target {
// 				count += right - left
// 				left++
// 			}
// 			if sum >= target {
// 				right--
// 			}
// 		}
// 	}
// 	return count
// }

// time O(n^2)
// space O(1) excluding output
// space O(n^2) worst case
