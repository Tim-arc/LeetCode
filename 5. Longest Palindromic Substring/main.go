package main

import (
	"bufio"
	"fmt"
	"os"
)

func expandAroundCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	longest := ""
	for i := 0; i < len(s); i++ {
		// Обработка нечетных палиндромов
		palindrome1 := expandAroundCenter(s, i, i)
		// Обработка четных палиндромов
		palindrome2 := expandAroundCenter(s, i, i+1)

		// Выбираем самый длинный палиндром
		if len(palindrome1) > len(longest) {
			longest = palindrome1
		}
		if len(palindrome2) > len(longest) {
			longest = palindrome2
		}
	}
	return longest
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s := ""
	fmt.Fscan(in, &s)
	fmt.Fprintln(out, longestPalindrome(s))
}
