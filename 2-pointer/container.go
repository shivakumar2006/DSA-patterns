// container with most water
// https://leetcode.com/problems/container-with-most-water/description/?envId=eabxy4hs&envType=problem-list-v2&utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(container(arr))
}

func container(arr []int) int {
	left := 0
	right := len(arr) - 1

	maxArea := 0

	for left < right {
		width := right - left
		height := min(arr[left], arr[right])
		area := width * height

		if arr[left] < arr[right] {
			left++
		} else {
			right--
		}

		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}
