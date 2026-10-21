// minimum window substring
// https://leetcode.com/problems/minimum-window-substring/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	required := len(need)
	formed := 0

	window := make(map[byte]int)

	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		window[ch]++

		if count, exists := need[ch]; exists && window[ch] == count {
			formed++
		}

		for formed == required {
			length := right - left + 1

			if length < minLen {
				minLen = length
				start = left
			}

			leftChar := s[left]
			window[leftChar]--

			if count, exists := need[leftChar]; exists && window[leftChar] < count {
				formed--
			}

			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
