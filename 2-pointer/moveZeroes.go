// move zeroes

package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(arr []int) {
	left := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] != 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
	}
}

// func moveZeroes(arr []int) {
// 	left := 0

// 	for right := 0; right < len(arr); right++ {
// 		if arr[right] != 0 {
// 			arr[left], arr[right] = arr[right], arr[left]
// 			left++
// 		}
// 	}
// }

// time O(n)
// space O(1)
