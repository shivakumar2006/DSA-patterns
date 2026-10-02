// 3 sum closest
// https://leetcode.com/problems/3sum-closest/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 2, 1, -4} // [-4, -1, 1, 2]
	target := 1
	res := closest(arr, target)
	fmt.Println(res)
}

func closest(arr []int, target int) int {
	sort.Ints(arr)

	closestSum := arr[0] + arr[1] + arr[2]

	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			diff := abs(sum - target)
			closestDiff := abs(closestSum - target)

			if diff < closestDiff {
				closestSum = sum
			}

			if sum == target {
				return sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closestSum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// func closest(arr []int, target int) int {
// 	sort.Ints(arr)

// 	// assume first 3 sum is the initial closest
// 	closestSum := arr[0] + arr[1] + arr[2]

// 	for i := 0; i < len(arr)-2; i++ {
// 		left := i + 1
// 		right := len(arr) - 1

// 		for left < right {
// 			sum := arr[i] + arr[left] + arr[right]
// 			diff := abs(sum - target)
// 			closestDiff := abs(closestSum - target)

// 			if diff < closestDiff {
// 				closestSum = sum
// 			}

// 			if sum == target {
// 				return sum
// 			}

// 			if sum < target {
// 				left++
// 			} else {
// 				right--
// 			}
// 		}
// 	}

// 	return closestSum
// }

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }
