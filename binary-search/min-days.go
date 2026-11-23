// 1482. Minimum Number of Days to Make m Bouquets
// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/description/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	bloomDay := []int{1, 10, 3, 10, 2} // 5 flowers
	m := 3
	k := 1
	fmt.Println(minDays(bloomDay, m, k))
}

func minDays(bloomDay []int, m, k int) int {
	left := 0
	right := 0

	for _, val := range bloomDay {
		left = min(left, val)   // 1
		right = max(right, val) // 10
	}

	for left < right {
		mid := left + (right-left)/2
		if canMakeBouquet(bloomDay, mid, m, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
} // time O(n log D) D -> no. of days
// space O(1)

func canMakeBouquet(bloomDay []int, mid, m, k int) bool {
	bouquet := 0
	flowers := 0

	for _, val := range bloomDay {
		if val <= mid {
			flowers += 1
		} else {
			flowers = 0
		}

		if flowers == k {
			bouquet += 1
			flowers = 0
		}
	}
	return bouquet >= m
} // time O(n), space O(1)

// overall time complexity is O(n) + O(n log D) = O(n log D)
// overall space complexity is O(1)

// make a m bouquet
// every bouquet have k adjacent flowers
// one flower only in one bouquet
