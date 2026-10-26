// max consecutive ones 3
// https://leetcode.com/problems/max-consecutive-ones-iii/description/

package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println(maxConsecutiveOnes(arr, k))
}

func maxConsecutiveOnes(arr []int, k int) int {
	left := 0
	zeroes := 0
	maxLen := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] == 0 {
			zeroes++
		}

		for zeroes > k {
			if arr[left] == 0 {
				zeroes--
			}
			left++
		}

		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

// func maxConsecutiveOnes(arr []int, k int) int {
// 	left := 0
// 	zeroes := 0
// 	maxLen := 0

// 	for right := 0; right < len(arr); right++ {
// 		if arr[right] == 0 {
// 			zeroes++
// 		}

// 		for zeroes > k {
// 			if arr[left] == 0 {
// 				zeroes--
// 			}
// 			left++
// 		}

// 		length := right - left + 1
// 		if length > maxLen {
// 			maxLen = length
// 		}
// 	}

// 	return maxLen
// }

// func maxConsecutiveOnes(arr []int, k int) int {
// 	left := 0
// 	zeroes := 0
// 	maxLen := 0

// 	for right := 0; right < len(arr); right++ {
// 		if arr[right] == 0 {
// 			zeroes++
// 		}

// 		if zeroes > k {
// 			if arr[left] == 0 {
// 				zeroes--
// 			}
// 			left++
// 		}

// 		length := right - left + 1

// 		if length > maxLen {
// 			maxLen = length
// 		}
// 	}

// 	return maxLen
// }

// func maxConsecutiveOnes(arr []int, k int) int {
// 	left := 0
// 	zeroes := 0
// 	maxLen := 0

// 	for right := 0; right < len(arr); right++ {
// 		if arr[right] == 0 {
// 			zeroes++
// 		}

// 		if zeroes > k {
// 			if arr[left] == 0 {
// 				zeroes--
// 			}
// 			left++
// 		}

// 		length := right - left + 1
// 		if length > maxLen {
// 			maxLen = length
// 		}
// 	}

// 	return maxLen
// }
