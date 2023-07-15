// remove-linked-list-elements
// Given the head of a linked list and an integer val,
// remove all the nodes of the linked list that has
// Node.val == val, and return the new head.
package main

import "log"

type example struct {
	nums []int
	val  int
}

func main() {
	examples := []example{
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6},
		{[]int{}, 1},
		{[]int{7, 7, 7, 7}, 7},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 2, 4, 5}, 2},
		{[]int{1, 2, 2, 2, 5}, 2},
		{[]int{1, 2, 2, 2, 2}, 2},
		{[]int{2, 2, 2, 2, 1}, 2},
	}

	for _, ex := range examples {
		log.Printf("input:  list = %+v, val = %d\n", ex.nums, ex.val)
		head := removeElements(fromSlice(ex.nums), ex.val)
		log.Printf("output: %+v", toSlice(head))
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	nHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := nHead
	for node := nHead.Next; node != nil; node = node.Next {
		if node.Val == val {
			prev.Next = node.Next
		} else {
			prev = node
		}
	}
	return nHead.Next
}

// ListNode defines a singly linked list.
// Provided by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
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

	return listNode
}

func toSlice(head *ListNode) []int {
	nodes := []int{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node.Val)
	}
	return nodes
}
