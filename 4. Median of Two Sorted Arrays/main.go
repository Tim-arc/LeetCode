package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	buff := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			buff = append(buff, nums1[i])
			i++
		} else {
			buff = append(buff, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		buff = append(buff, nums1[i])
		i++
	}

	for j < len(nums2) {
		buff = append(buff, nums2[j])
		j++
	}

	n := len(buff)
	if n%2 == 1 {
		return float64(buff[n/2])
	} else {
		return float64(buff[n/2-1]+buff[n/2]) / 2.0
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
