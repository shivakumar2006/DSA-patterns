// Maximum subarray sum with one deletion
// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, -2, 0, 3}
	fmt.Println(oneDeleteion(arr))
}

func oneDeleteion(arr []int) int {
	noDelete := arr[0]
	oneDelete := math.MinInt
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		prevNoDelete := noDelete
		prevOneDelete := oneDelete

		noDelete = max(arr[i], prevNoDelete+arr[i])

		if prevOneDelete == math.MinInt {
			oneDelete = prevNoDelete
		} else {
			oneDelete = max(prevNoDelete, prevOneDelete+arr[i])
		}

		result = max(result, max(noDelete, oneDelete))
	}

	return result
}
