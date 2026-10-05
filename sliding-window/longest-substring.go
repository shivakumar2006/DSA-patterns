// Longest Substring with K units
// https://www.geeksforgeeks.org/problems/longest-k-unique-characters-substring0853/1

package main

import "fmt"

func main() {
	s := "aabacbebebe"
	k := 3
	fmt.Println(lengthOfLongestSubstringKUnique(s, k))
}

func lengthOfLongestSubstringKUnique(s string, k int) int {
	freq := make(map[byte]int)

	left := 0
	result := -1

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for len(freq) > k {
			freq[s[left]]--

			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}

			left++
		}

		if len(freq) == k {
			length := right - left + 1

			if length > result {
				result = length
			}
		}
	}

	return result
}

// time O(n)
// space O(1)
