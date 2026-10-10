package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 1, 1, 0, 0, 0, 1, 0}
	fmt.Println(rearrange(arr))
}

func rearrange(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		if arr[left] == 0 {
			left++
		} else if arr[right] == 1 {
			right--
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return arr
}

// func rearrange(arr []int) []int {
// 	i := 0
// 	j := len(arr) - 1

// 	for i < j {
// 		if arr[i] == 0 {
// 			i++
// 		} else if arr[j] == 1 {
// 			j--
// 		} else {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 			j--
// 		}
// 	}

// 	return arr
// }

// func rearrange(arr []int) []int {
// 	i := 0
// 	j := len(arr) - 1

// 	for i < j {
// 		if arr[i] == 0 {
// 			i++
// 		}
// 		if arr[j] == 1 {
// 			j--
// 		}
// 		if arr[i] == 1 && arr[j] == 0 {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 			j--
// 		}
// 	}
// 	return arr
// }

// time O(n)
// space O(1)
// in place swapping
