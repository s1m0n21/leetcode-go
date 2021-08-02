/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/min-stack/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/3 12:14 上午
 */

package _min_stack

type MinStack struct {
	min   int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   1<<63 - 1,
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		this.min = val
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	n := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if n == this.min {
		this.min = 1<<63 - 1
		for _, v := range this.stack {
			this.min = func(a, b int) int {
				if a < b {
					return a
				}
				return b
			}(this.min, v)
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
