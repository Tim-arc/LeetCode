package main

import (
	"fmt"
)

func getInt(s rune) int {
	if s == 'I' {
		return 1
	} else if s == 'V' {
		return 5
	} else if s == 'X' {
		return 10
	} else if s == 'L' {
		return 50
	} else if s == 'C' {
		return 100
	} else if s == 'D' {
		return 500
	} else if s == 'M' {
		return 1000
	} else {
		return 0
	}
}

func romanToInt(s string) int {
	result := 0
	s_map := make(map[int]rune)
	for i, el := range s {
		s_map[i] = el
	}

	for i := 0; i < len(s_map); i++ {
		if getInt(s_map[i]) >= getInt(s_map[i+1]) {
			result += getInt(s_map[i])
		} else {
			result += getInt(s_map[i+1]) - getInt(s_map[i])
			i++
		}
	}

	return result

}

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}
