// https://leetcode.com/problems/reverse-linked-list/
// Given the head of a singly linked list,
// reverse the list, and return the reversed list.
package main

import "log"

type example struct {
	nums []int
}

func main() {
	examples := []example{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 2}},
		{[]int{}},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v\n", ex.nums)
		reversed := reverseList(fromSlice(ex.nums))
		log.Printf("output: %+v\n", toSlice(reversed))
	}
}

// ListNode defines a singly linked list.
// Provided by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://www.interviewbit.com/blog/reverse-a-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return res
}

func fromSlice(digits []int) *ListNode {
	var listNode *ListNode
	for i := len(digits) - 1; i >= 0; i-- {
		if listNode == nil {
			listNode = &ListNode{
				Val:  digits[i],
				Next: nil,
			}
		} else {
			tmp := &ListNode{
				Val:  digits[i],
				Next: listNode,
			}
			listNode = tmp
		}
	}
	/*
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
	*/
	return listNode
}

func toSlice(head *ListNode) []int {
	nodes := []int{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node.Val)
	}
	return nodes
}
