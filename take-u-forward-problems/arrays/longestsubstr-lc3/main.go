package main

import (
	"fmt"
)

// with hash map
func lengthOfLongestSubstring(s string) int {
	maxLen, count, lastIndex := 0, 0, -1
	seen := map[byte]int{}

	for i := 0; i < len(s); i++ {
		if index, ok := seen[s[i]]; ok && index > lastIndex {
			count = i - (index + 1)
			lastIndex = index
		}

		seen[s[i]] = i
		count++

		if maxLen < count {
			maxLen = count
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwwke"))
	fmt.Println(lengthOfLongestSubstring("dvdfedsgvh"))
}
