// Fruits into basket
// https://leetcode.com/problems/fruit-into-baskets/description/

package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 2}
	fmt.Println(basket(arr))
}

func basket(arr []int) int {
	freq := make(map[int]int)

	left := 0
	result := -1

	for right := 0; right < len(arr); right++ {
		freq[arr[right]]++

		for len(freq) > 2 {
			freq[arr[left]]--

			if freq[arr[left]] == 0 {
				delete(freq, arr[left])
			}

			left++
		}

		length := right - left + 1

		if length > result {
			result = length
		}
	}

	return result
}

// func basket(arr []int) int {
// 	freq := make(map[int]int)

// 	left := 0
// 	result := -1

// 	for right := 0; right < len(arr); right++ {
// 		freq[arr[right]]++

// 		for len(freq) > 2 {
// 			freq[arr[left]]--

// 			if freq[arr[left]] == 0 {
// 				delete(freq, arr[left])
// 			}

// 			left++
// 		}

// 		length := right - left + 1
// 		if length > result {
// 			result = length
// 		}
// 	}

// 	return result
// }

// time O(n)
// space O(n)
