// grumpy bookstore owner
// https://leetcode.com/problems/grumpy-bookstore-owner/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	fmt.Println(maxSatisfied(customers, grumpy, minutes))
}

func maxSatisfied(customers, grumpy []int, minutes int) int {
	totalSatisfied := 0

	// count the customers already satisfied
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			totalSatisfied += customers[i]
		}
	}

	// calculate extra customers saved
	windowSum := 0
	maxExtra := 0

	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			windowSum += customers[i]
		}
	}

	maxExtra = windowSum

	// slide fixed sizezd window
	for right := minutes; right < len(customers); right++ {
		if grumpy[right] == 1 {
			windowSum += customers[right]
		}

		left := right - minutes
		if grumpy[left] == 1 {
			windowSum -= customers[left]
		}

		if windowSum > maxExtra {
			maxExtra = windowSum
		}
	}

	return totalSatisfied + maxExtra
}

// time O(n)
// space O(1)
