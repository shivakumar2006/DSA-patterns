// longest repeating characters replacements
// https://leetcode.com/problems/longest-repeating-character-replacement/description/

package main

import "fmt"

func main() {
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)

	left := 0
	maxFreq := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		if freq[s[right]] > maxFreq {
			maxFreq = freq[s[right]]
		}

		length := right - left + 1

		if length-maxFreq > k {
			freq[s[left]]--
			left++
		}

		length = right - left + 1

		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

// time O(n)
// space O(1)
