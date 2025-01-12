package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(str) < len(prefix) || str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}
