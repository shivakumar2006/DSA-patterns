package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(sum(arr))
}

func sum(arr []int) [][]int {
	sort.Ints(arr)

	result := [][]int{}

	// it runs till len(arr)-2 because, we need 3 values for answer if it is go on till len(arr)-1 then we have left only 2 values but we want 2 values in our output
	for i := 0; i < len(arr)-2; i++ {
		// skip same fixed element
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]

			if sum == 0 {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				left++
				right--

				// duplicate left value skip
				for left < right && arr[left] == arr[left-1] {
					left++
				}

				// duplicate right value skip
				for left < right && arr[right] == arr[right+1] {
					right--
				}

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// time O(n^2)
// space O(1) excluding output
// space O(n^2) worst case
