/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/implement-stack-using-queues/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/1 5:18 下午
 */

package _implement_stack_using_queues

type MyStack struct {
	queue []int
}

// Constructor Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{make([]int, 0)}
}

// Push /** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}

// Pop Removes the element on top of the stack and returns that element.
func (this *MyStack) Pop() int {
	n := this.Top()
	this.queue = this.queue[:len(this.queue)-1]
	return n
}

// Top Get the top element.
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

// Empty Returns whether the stack is empty.
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
