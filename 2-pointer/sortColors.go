// sort colors
// https://leetcode.com/problems/sort-colors/description/

package main

import "fmt"

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}

func sortColors(arr []int) {
	low := 0
	mid := 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

// time O(n), space O(1)
