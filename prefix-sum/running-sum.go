// running sum of 1d array
// https://leetcode.com/problems/running-sum-of-1d-array/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(runningSum(arr))
}

func runningSum(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	return arr
}

// time O(n)
// space O(1)
