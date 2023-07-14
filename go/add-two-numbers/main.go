// https://leetcode.com/problems/add-two-numbers/
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
package main

import "log"

type example struct {
	l1Slice []int
	l2Slice []int
}

func main() {
	examples := []example{
		{[]int{2, 4, 3}, []int{5, 6, 4}},
		{[]int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
	}

	for _, example := range examples {
		out := addTwoNumbers(fromSlice(example.l1Slice), fromSlice(example.l2Slice))

		for node := out; node != nil; node = node.Next {
			log.Println(node.Val)
		}
	}
}

// ListNode represents a singly linked list of integers.
// Provided by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	l1Node := l1
	l2Node := l2
	carry := false

	c, carry := add(l1Node, l2Node, carry)

	var l1Next *ListNode
	l1Next = nil
	if l1Node != nil {
		l1Next = l1Node.Next
	}

	var l2Next *ListNode
	l2Next = nil
	if l2Node != nil {
		l2Next = l2Node.Next
	}

	// Move the carry digit forward, preferring the first list.
	// If the first list is empty, we'll carry forward on the second list.
	// When both lists are exhausted, create a new node for the carry digit.
	if carry {
		if l1Next != nil {
			l1Next.Val++
		} else if l2Next != nil {
			l2Next.Val++
		} else {
			l1Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	}

	return &ListNode{
		Val:  c,
		Next: addTwoNumbers(l1Next, l2Next),
	}
}

func add(l1 *ListNode, l2 *ListNode, doCarry bool) (int, bool) {
	a := 0
	if l1 != nil {
		a = l1.Val
	}

	b := 0
	if l2 != nil {
		b = l2.Val
	}

	c := a + b

	if doCarry {
		c++
	}

	// Directions tell us that 0 <= digit <= 10,
	// so we know we're working in base-10.
	carry := false
	if c >= 10 {
		carry = true
	}

	return c % 10, carry
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
