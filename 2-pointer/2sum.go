package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	res := sum(arr, target)
	fmt.Println(res)
}

func sum(arr []int, tar int) []int {
	i := 0
	j := len(arr) - 1

	for i < j {
		sum := arr[i] + arr[j]
		if sum == tar {
			return []int{i + 1, j + 1}
		}
		if sum > tar {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}
