// valid palindrome or not

package main

import "fmt"

func main() {
	s := "racecar"
	fmt.Println(valid(s))
}

func valid(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// func valid(s string) bool {
// 	left := 0
// 	right := len(s) - 1

// 	for left < right {
// 		if s[left] != s[right] {
// 			return false
// 		}

// 		left++
// 		right--
// 	}

// 	return true
// }

// time : O(n)
// space : O(1)
