// square of the sorted array
//https://leetcode.com/problems/squares-of-a-sorted-array/description/

package main

import "fmt"

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortSquares(arr))
}

func sortSquares(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	left := 0
	right := n - 1

	for pos := n - 1; pos >= 0; pos-- {
		if abs(arr[left]) > abs(arr[right]) {
			result[pos] = arr[left] * arr[left]
			left++
		} else {
			result[pos] = arr[right] * arr[right]
			right--
		}
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// func sortSquares(arr []int) []int {
// 	n := len(arr)
// 	result := make([]int, n)

// 	left := 0
// 	right := n - 1

// 	for pos := n - 1; pos >= 0; pos-- {
// 		if abs(arr[left]) > abs(arr[right]) {
// 			result[pos] = arr[left] * arr[left]
// 			left++
// 		} else {
// 			result[pos] = arr[right] * arr[right]
// 			right--
// 		}
// 	}
// 	return result
// }

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// func sortSquares(arr []int) []int {
// 	n := len(arr)
// 	result := make([]int, n)

// 	left := 0
// 	right := n - 1

// 	for pos := n - 1; pos >= 0; pos-- {
// 		if abs(arr[left]) > abs(arr[right]) {
// 			result[pos] = arr[left] * arr[left]
// 			left++
// 		} else {
// 			result[pos] = arr[right] * arr[right]
// 			right--
// 		}
// 	}
// 	return result
// }

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// func sortSquares(arr []int) []int {
// 	n := len(arr)
// 	result := make([]int, n)

// 	left := 0
// 	right := n - 1

// 	for pos := n - 1; pos >= 0; pos-- {
// 		if abs(arr[left]) > abs(arr[right]) {
// 			result[pos] = arr[left] * arr[left]
// 			left++
// 		} else {
// 			result[pos] = arr[right] * arr[right]
// 			right--
// 		}
// 	}
// 	return result
// }

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }
