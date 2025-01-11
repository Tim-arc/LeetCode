package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var current *ListNode

	sum1 := 0
	carry := 0

	for l1 != nil || l2 != nil {
		sum1 = carry
		if l1 != nil {
			sum1 += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum1 += l2.Val
			l2 = l2.Next
		}
		carry = sum1 / 10
		sum1 %= 10
		if result == nil {
			result = &ListNode{Val: sum1}
			current = result
		} else {
			current.Next = &ListNode{Val: sum1}
			current = current.Next
		}
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}

	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	result := addTwoNumbers(l1, l2)

	for node := result; node != nil; node = node.Next {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
