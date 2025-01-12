package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, char := range s {
		if pos, exists := charMap[char]; exists && pos >= start {
			start = pos + 1
		}
		charMap[char] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func main() {
	str := "abcabcbb"
	result := lengthOfLongestSubstring(str)
	fmt.Println(result)

	str = "pwwkew"
	result = lengthOfLongestSubstring(str)
	fmt.Println(result)
}
