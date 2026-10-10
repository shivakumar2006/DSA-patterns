package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	res := sum(arr, target)
	fmt.Println(res)
}

func sum(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

// func sum(arr []int, target int) []int {
// 	i := 0
// 	j := len(arr) - 1

// 	for i < j {
// 		sum := arr[i] + arr[j]

// 		if sum == target {
// 			return []int{i + 1, j + 1}
// 		}
// 		if sum > target {
// 			j--
// 		} else {
// 			i++
// 		}
// 	}
// 	return []int{-1, -1}
// }

// func sum(arr []int, tar int) []int {
// 	i := 0
// 	j := len(arr) - 1

// 	for i < j {
// 		sum := arr[i] + arr[j]
// 		if sum == tar {
// 			return []int{i + 1, j + 1}
// 		}
// 		if sum > tar {
// 			j--
// 		} else {
// 			i++
// 		}
// 	}
// 	return []int{-1, -1}
// }

// time complexity O(n)
// space O(1)
