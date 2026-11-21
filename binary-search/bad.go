// first bad version
// https://leetcode.com/problems/first-bad-version/description/?utm_source=chatgpt.com

package main

import "fmt"

var bad int

func main() {
	n := 5
	bad = 4

	fmt.Println(firstBadVersion(n))
}

func isBadVersion(version int) bool {
	return version >= bad
} // time O(1), space O(1)

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
} // time O(log n), space O(1)

// O(log n) + O(1) = O(log n)

// time O(log n)
// space O(1)
