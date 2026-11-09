// Max product subarray
// https://leetcode.com/problems/maximum-product-subarray/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	arr := []int{2, 3, -2, 4}
	fmt.Println(maxProdSubarray(arr))
}

func maxProdSubarray(arr []int) int {
	maxProd := arr[0]
	minProd := arr[0]
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(arr[i], maxProd+arr[i])
		minProd = min(arr[i], minProd+arr[i])

		result = max(result, maxProd)
	}
	return result
}

// func maxProdSubarray(arr []int) int {
// 	maxProd := arr[0]
// 	minProd := arr[0]
// 	result := arr[0]

// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < 0 {
// 			maxProd, minProd = minProd, maxProd
// 		}

// 		maxProd = max(arr[i], maxProd*arr[i])
// 		minProd = min(arr[i], minProd*arr[i])

// 		result = max(result, maxProd)
// 	}
// 	return result
// }

// func maxProdSubarray(arr []int) int {
// 	maxProd := arr[0]
// 	minProd := arr[0]
// 	result := arr[0]

// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < 0 {
// 			maxProd, minProd = minProd, maxProd
// 		}

// 		maxProd = max(arr[i], maxProd*arr[i])
// 		minProd = min(arr[i], minProd*arr[i])

// 		result = max(result, maxProd)
// 	}
// 	return result
// }

// func maxProdSubarray(arr []int) int {
// 	maxProd := arr[0]
// 	minProd := arr[0]
// 	result := arr[0]

// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < 0 {
// 			maxProd, minProd = minProd, maxProd
// 		}

// 		maxProd = max(arr[i], maxProd*arr[i])
// 		minProd = min(arr[i], minProd*arr[i])

// 		result = max(result, maxProd)
// 	}

// 	return result
// }

// time O(n)
// space O(1)
