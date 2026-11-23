// 1283. Find the Smallest Divisor Given a Threshold
// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 9}
	thresHold := 6
	fmt.Println(divisor(arr, thresHold))
}

func divisor(arr []int, thresHold int) int {
	left := 1
	right := 0

	for _, val := range arr {
		right = max(right, val)
	}

	for left < right {
		mid := left + (right-left)/2
		if canDivide(arr, mid, thresHold) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
} // time O(log m), space O(1)

func canDivide(arr []int, divisor, thresHold int) bool {
	result := 0
	for _, val := range arr {
		result += (val + divisor - 1) / divisor
	}
	return result <= thresHold
} // time O(n), space O(1)

// time O(n log m)
// space O(1)

// left := 1
// right := max of array which is 9
// now find mid (1+9)/2 = 5 (divisor)
// canDivide check (1+5-1)/5 + (2+5-1)/5 + (5+5-1)/5 + (9+5-1)/5 = 1 + 1 + 1 + 2 = 5
// 5 <= 6 true so now right = mid which is 5

// new mid (1+5)/2 = 3 (divisor)
// canDivide check (1+3-1)/3 + (2+3-1)/3 + (5+3-1)/3 + (9+3-1)/3 = 1 + 1 + 2 + 3 = 7
// 7 <= 6 false so now left = mid + 1 which is 4

// so the answer is 5
