package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	buff := make([]int, 0)
	for x > 0 {
		buff = append(buff, x%10)
		x /= 10
	}
	for i := 0; i < int(len(buff)/2); i++ {
		if buff[i] != buff[len(buff)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	input := 10022201

	fmt.Println(isPalindrome(input))

}
