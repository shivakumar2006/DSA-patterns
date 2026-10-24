// remove element
// https://leetcode.com/problems/remove-element/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(element(arr, val))
}

func element(arr []int, val int) int {
	left := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] != val {
			arr[left] = arr[right]
			left++
		}
	}

	return left
}

// func element(arr []int, val int) int {
// 	slow := 0

// 	for fast := 0; fast < len(arr); fast++ {
// 		if arr[fast] != val {
// 			arr[slow] = arr[fast]
// 			slow++
// 		}
// 	}

// 	return slow
// }

// func element(arr []int, value int) int {
// 	slow := 0

// 	for fast := 0; fast < len(arr); fast++ {
// 		if arr[fast] != value {
// 			arr[slow] = arr[fast]
// 			slow++
// 		}
// 	}

// 	return slow
// }

// func element(arr []int, value int) int {
// 	slow := 0

// 	for fast := 0; fast < len(arr); fast++ {
// 		if arr[fast] != value {
// 			arr[slow] = arr[fast]
// 			slow++
// 		}
// 	}

// 	return slow
// }
