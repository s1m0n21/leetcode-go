/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/three-in-one-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/23 2:23 下午
 */

package interview_03_01_three_in_one_lcci

type TripleInOne struct {
	stacks [3][]int
}

func Constructor(stackSize int) TripleInOne {
	stacks := [3][]int{}
	for i := range stacks {
		stacks[i] = make([]int, 0, stackSize)
	}
	return TripleInOne{stacks}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if len(this.stacks[stackNum]) == cap(this.stacks[stackNum]) {
		return
	}

	this.stacks[stackNum] = append(this.stacks[stackNum], value)
}

func (this *TripleInOne) Pop(stackNum int) int {
	n := len(this.stacks[stackNum])
	if n == 0 {
		return -1
	}

	v := this.stacks[stackNum][n-1]
	this.stacks[stackNum] = this.stacks[stackNum][:n-1]

	return v
}

func (this *TripleInOne) Peek(stackNum int) int {
	n := len(this.stacks[stackNum])
	if n == 0 {
		return -1
	}

	return this.stacks[stackNum][n-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.stacks[stackNum]) == 0
}
