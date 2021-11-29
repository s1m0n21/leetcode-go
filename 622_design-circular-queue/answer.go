/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/design-circular-queue/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/29 9:47 上午
 */

package _design_circular_queue

type MyCircularQueue struct {
	array    []int
	maxSize  int
	currSize int
	head     int
	tail     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		array:   make([]int, k),
		maxSize: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.currSize == this.maxSize {
		return false
	}
	this.array[this.tail] = value
	this.tail++
	this.tail %= this.maxSize
	this.currSize++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.currSize == 0 {
		return false
	}
	this.head++
	this.head %= this.maxSize
	this.currSize--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.currSize == 0 {
		return -1
	}
	return this.array[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.currSize == 0 {
		return -1
	}
	ptr := this.tail
	ptr--
	ptr = (ptr + this.maxSize) % this.maxSize
	return this.array[ptr]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.currSize == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.maxSize == this.currSize
}
