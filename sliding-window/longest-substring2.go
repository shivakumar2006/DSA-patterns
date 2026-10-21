// longest substring without repeating characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for freq[s[right]] > 1 {
			freq[s[left]]--
			left++
		}

		length := right - left + 1

		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

// func lengthOfLongestSubstring(s string) int {
// 	freq := make(map[byte]int)

// 	left := 0
// 	maxLen := 0

// 	for right := 0; right < len(s); right++ {
// 		freq[s[right]]++

// 		for freq[s[right]] > 1 {
// 			freq[s[left]]--
// 			left++
// 		}

// 		length := right - left + 1

// 		if length > maxLen {
// 			maxLen = length
// 		}
// 	}

// 	return maxLen
// }
