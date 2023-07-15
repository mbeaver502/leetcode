// https://leetcode.com/problems/implement-stack-using-queues/
// Implement a last-in-first-out (LIFO) stack using only two queues.
// The implemented stack should support all the functions of a
// normal stack (push, top, pop, and empty).
package main

import "log"

/*
Your MyStack object will be instantiated and called as such:
obj := Constructor();
obj.Push(x);
param_2 := obj.Pop();
param_3 := obj.Top();
param_4 := obj.Empty();
*/
func main() {
	s := Constructor()
	s.Push(2)
	s.Push(2)
	log.Printf("top: %d\n", s.Top())
	log.Printf("pop: %d\n", s.Pop())
	log.Printf("empty: %+v\n", s.Empty())
}

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{
		stack: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	val := this.Top()
	this.stack = this.stack[:len(this.stack)-1]
	return val
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}
