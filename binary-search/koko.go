// koko eating banana
// https://leetcode.com/problems/koko-eating-bananas/description/?utm_source=chatgpt.com

package main

import (
	"fmt"
	"slices"
)

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(koko(piles, h))
}

func koko(arr []int, h int) int {
	start := 1
	end := slices.Max(arr)
	for start < end {
		mid := start + (end-start)/2
		if finish(arr, mid, h) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func finish(arr []int, speed, h int) bool {
	totalHours := 0
	for _, i := range arr {
		totalHours = totalHours + (i+speed-1)/speed
	}
	return totalHours <= h
}
