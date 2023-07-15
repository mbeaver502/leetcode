// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Given the head of a sorted linked list, delete all duplicates
// such that each element appears only once.
// Return the linked list sorted as well.
package main

import "log"

// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.
type example struct {
	digits []int
}

func main() {
	examples := []example{
		{[]int{1, 1, 2}},
		{[]int{1, 1, 2, 3, 3}},
		// Test cases from https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/comments/1895135
		{[]int{1, 1, 1, 1}},
		{[]int{1, 2, 3, 4, 4}},
		{[]int{1}},
		{[]int{1, 1, 1, 1, 1, 2, 3, 3, 3, 4, 4}},
		{[]int{-99, -30, -30, -30, 1, 1, 1, 2, 2, 2, 2, 2, 33, 33, 33, 33}},
		{[]int{}},
	}

	for _, ex := range examples {
		log.Printf("input:  %v", ex.digits)
		l := deleteDuplicates(fromSlice(ex.digits))
		log.Printf("output: %v", toSlice(l))
	}
}

// ListNode defines a singly-linked list.
// Provided by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lastNode := head
	for node := head.Next; node != nil; node = node.Next {
		if node.Val == lastNode.Val {
			lastNode.Next = node.Next
			continue
		}
		lastNode = node
	}

	return head
}

func fromSlice(digits []int) *ListNode {
	var listNode *ListNode
	for _, d := range digits {
		if listNode == nil {
			listNode = &ListNode{
				Val:  d,
				Next: nil,
			}
		} else {
			tmp := &ListNode{
				Val:  d,
				Next: listNode,
			}
			listNode = tmp
		}
	}
	return listNode
}

func toSlice(head *ListNode) []int {
	list := []int{}
	for node := head; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	out := make([]int, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		out[len(list)-i-1] = list[i]
	}
	return out
}
