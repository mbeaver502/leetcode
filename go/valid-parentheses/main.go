// https://leetcode.com/problems/valid-parentheses/
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
// An input string is valid if:
//   - Open brackets must be closed by the same type of brackets.
//   - Open brackets must be closed in the correct order.
//   - Every close bracket has a corresponding open bracket of the same type.
package main

import (
	"log"
)

type example struct {
	value string
}

func main() {
	examples := []example{
		{"()"},                           // valid
		{"()[]{}"},                       // valid
		{"(]"},                           // false
		{"((({{[{[[{(([{}]))}]]}]}})))"}, // valid
		{"((({{[{[[{((){}]))}]]}]}})))"}, // invalid
		{"([)]"},                         // invalid
		{"{[]}"},                         // valid
	}

	for _, example := range examples {
		log.Printf("input = %s", example.value)
		log.Printf("output = %v", isValid(example.value))
	}
}

func isValid(s string) bool {
	// We'll simulate a stack using a slice.
	stack := []rune{}

	// We're guaranteed that s consists of parentheses only '()[]{}'.
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, '(')
		case ')':
			if len(stack) < 1 {
				return false
			}
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]

		case '[':
			stack = append(stack, '[')
		case ']':
			if len(stack) < 1 {
				return false
			}
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]

		case '{':
			stack = append(stack, '{')
		case '}':
			if len(stack) < 1 {
				return false
			}
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}
