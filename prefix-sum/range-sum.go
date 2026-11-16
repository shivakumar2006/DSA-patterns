// Range sum query
// https://leetcode.com/problems/range-sum-query-immutable/

package main

import "fmt"

type NumArray struct {
	Prefix []int
}

func Constructor(arr []int) NumArray {
	prefix := make([]int, len(arr))

	prefix[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	return NumArray{Prefix: prefix}
}

func (this *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return this.Prefix[right]
	}

	return this.Prefix[right] - this.Prefix[left-1]
}

func main() {
	arr := []int{-2, 0, 3, -5, 2, -1}

	numArray := Constructor(arr)

	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
