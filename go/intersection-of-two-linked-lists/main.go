// https://leetcode.com/problems/intersection-of-two-linked-lists/
// Given the heads of two singly linked-lists headA and headB,
// return the node at which the two lists intersect.
// If the two linked lists have no intersection at all, return null.
package main

import "log"

func main() {
	intersect := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: intersect,
		},
	}
	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: intersect,
			},
		},
	}

	intersectNode := getIntersectionNode(headA, headB)
	if intersectNode != nil {
		log.Printf("output: node = %+v, val = %d, next = %+v", intersectNode, intersectNode.Val, intersectNode.Next)
	} else {
		log.Printf("output: node = %+v", intersectNode)
	}

	intersect = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	headA = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  1,
				Next: intersect,
			},
		},
	}
	headB = &ListNode{
		Val:  3,
		Next: intersect,
	}

	intersectNode = getIntersectionNode(headA, headB)
	if intersectNode != nil {
		log.Printf("output: node = %+v, val = %d, next = %+v", intersectNode, intersectNode.Val, intersectNode.Next)
	} else {
		log.Printf("output: node = %+v", intersectNode)
	}

	headA = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	headB = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}

	intersectNode = getIntersectionNode(headA, headB)
	if intersectNode != nil {
		log.Printf("output: node = %+v, val = %d, next = %+v", intersectNode, intersectNode.Val, intersectNode.Next)
	} else {
		log.Printf("output: node = %+v", intersectNode)
	}

	headA = &ListNode{
		Val:  1,
		Next: nil,
	}
	headB = &ListNode{
		Val:  2,
		Next: nil,
	}

	intersectNode = getIntersectionNode(headA, headB)
	if intersectNode != nil {
		log.Printf("output: node = %+v, val = %d, next = %+v", intersectNode, intersectNode.Val, intersectNode.Next)
	} else {
		log.Printf("output: node = %+v", intersectNode)
	}
}

// ListNode defines a singly linked list.
// Provided by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newA := headA
	newB := headB

	lenA := getLen(newA)
	lenB := getLen(newB)
	lenDiff := lenA - lenB
	maxLen := lenA

	// Let's adjust the starting point of each list
	// so that they start from the same relative position.
	if lenDiff > 0 {
		newA = advanceBy(newA, lenDiff)
	} else if lenDiff < 0 {
		maxLen = lenB
		lenDiff *= (-1)
		newB = advanceBy(newB, lenDiff)
	}

	// Let's walk along each list in parallel
	// until we find the same node in both lists.
	for i := lenDiff; i < maxLen; i++ {
		if newA == nil || newB == nil {
			break
		}
		if newA == newB {
			return newA
		}
		newA = newA.Next
		newB = newB.Next
	}

	return nil
}

func getLen(head *ListNode) int {
	result := 0
	for node := head; node != nil; node = node.Next {
		result++
	}
	return result
}

func advanceBy(head *ListNode, by int) *ListNode {
	for i := 0; i < by; i++ {
		head = head.Next
	}
	return head
}
