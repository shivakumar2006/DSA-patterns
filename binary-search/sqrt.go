// sqrt(x)
// https://leetcode.com/problems/sqrtx/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	x := 4
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	left := 1
	right := x
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		sqrt := mid * mid
		if sqrt > x {
			right = mid - 1
		} else {
			result = mid
			left = mid + 1
		}
	}
	return result
} // Time: O(log n), Space: O(1)
