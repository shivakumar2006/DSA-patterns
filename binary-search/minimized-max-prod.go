// 2064. Minimized Maximum of Products Distributed to Any Store
// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/description/?utm_source=chatgpt.com

package main

import (
	"fmt"
)

func main() {
	n := 6
	quantities := []int{11, 6}

	fmt.Println(findProd(n, quantities))
}

func findProd(n int, quantities []int) int {
	left := 1
	right := 0

	for _, q := range quantities {
		right = max(right, q)
	}

	for left < right {
		mid := left + (right-left)/2
		if check(n, quantities, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
} // time complexity O(n log m)
// space O(1)

func check(n int, quantities []int, max int) bool {
	count := 0

	for _, q := range quantities {
		count += (q + max - 1) / max
	}

	return count <= n
} // time O(n), space O(1)

// overall time O(n log m) space O(1)
