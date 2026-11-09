// reverse string
// https://leetcode.com/problems/reverse-string/?utm_source=chatgpt.com

package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(reverse([]byte(s)))
}

func reverse(s []byte) string {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}

// func reverse(s []byte) string {
// 	left := 0
// 	right := len(s) - 1
// 	for left < right {
// 		s[left], s[right] = s[right], s[left]
// 		left++
// 		right--
// 	}

// 	return string(s)
// }

// time O(n)
// space O(1)

// func main() {
// 	s := []byte{'h', 'e', 'l', 'l', 'o'}
// 	reverse(s)
// 	fmt.Println(string(s))
// }

// func reverse(s []byte) {
// 	left := 0
// 	right := len(s) - 1

// 	for left < right {
// 		s[left], s[right] = s[right], s[left]
// 		left++
// 		right--
// 	}
// }
