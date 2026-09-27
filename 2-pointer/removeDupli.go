package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 2, 2, 3, 3, 3, 4}
	res := removeDupli(arr)
	fmt.Println(res)
}

func removeDupli(arr []int) int {
	i := 0
	j := i + 1

	res := 1

	for j < len(arr) {
		if arr[i] == arr[j] {
			j++
			continue
		}
		i++
		res++
		arr[i] = arr[j]
	}
	return res
}
