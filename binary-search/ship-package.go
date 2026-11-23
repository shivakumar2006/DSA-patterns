// Capacity to ship packages within D days
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5

	fmt.Println(shipWithinDays(weights, days))
}

func shipWithinDays(weights []int, days int) int {
	left := 0
	right := 0

	for _, weight := range weights {
		left = max(left, weight)
		right += weight
	}

	for left < right {
		mid := left + (right-left)/2
		if canShip(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
} // time O(n log s), space O(1)
// s = sum of all package weights
// n = number of packages

func canShip(weights []int, capacity, days int) bool {
	daysNeeded := 1
	currentWeight := 0

	for _, weight := range weights {
		if currentWeight+weight > capacity {
			daysNeeded++
			currentWeight = 0
		}

		currentWeight += weight
	}
	return daysNeeded <= days
} // time O(n), space O(1)

// same as split array problem
// we are searching for the minimum capacity that can ship all packages within D days
// so left = max element and right = sum of all elements
// and we use binary search to find the minimum capacity
// if canShip returns true then right = mid
// else left = mid + 1
