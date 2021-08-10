/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/implement-queue-using-stacks/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/11 2:17 上午
 */

package _implement_queue_using_stacks

type MyQueue struct {
	slice []int
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

// Push Push element x to the back of queue.
func (this *MyQueue) Push(x int)  {
	this.slice = append(this.slice, x)
}

// Pop Removes the element from in front of queue and returns that element.
func (this *MyQueue) Pop() int {
	defer func() {
		this.slice = this.slice[1:]
	}()
	return this.Peek()
}

// Peek Get the front element.
func (this *MyQueue) Peek() int {
	return this.slice[0]
}

// Empty Returns whether the queue is empty.
func (this *MyQueue) Empty() bool {
	return len(this.slice) == 0
}
