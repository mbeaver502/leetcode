// https://leetcode.com/problems/linked-list-cycle/
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer. Internally, pos is used to denote
// the index of the node that tail's next pointer is connected to.
// Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.
package main

import "log"

func main() {
	three := &ListNode{
		Val:  3,
		Next: nil,
	}
	two := &ListNode{
		Val:  2,
		Next: nil,
	}
	zero := &ListNode{
		Val:  0,
		Next: nil,
	}
	negFour := &ListNode{
		Val:  -4,
		Next: nil,
	}

	three.Next = two
	two.Next = zero
	zero.Next = negFour
	negFour.Next = two

	head := three
	log.Printf("output: %+v", hasCycle(head))

	one := &ListNode{
		Val:  1,
		Next: nil,
	}
	two = &ListNode{
		Val:  2,
		Next: nil,
	}
	one.Next = two

	head = one
	log.Printf("output: %+v", hasCycle(head))

	one = &ListNode{
		Val:  1,
		Next: nil,
	}
	two = &ListNode{
		Val:  2,
		Next: one,
	}
	one.Next = two

	head = one
	log.Printf("output: %+v", hasCycle(head))

	head = &ListNode{
		Val:  1,
		Next: nil,
	}
	log.Printf("output: %+v", hasCycle(head))
}

// ListNode defines a singly linked list.
// Provided by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	for node := head; node != nil; node = node.Next {
		if _, found := seen[node]; found {
			return true
		}
		seen[node] = true
	}
	return false
}
