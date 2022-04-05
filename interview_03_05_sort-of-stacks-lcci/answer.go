/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/sort-of-stacks-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/4 17:25
 */

package interview_03_05_sort_of_stacks_lcci

import "sort"

type SortedStack struct {
	stack []int
}

func Constructor() SortedStack {
	return SortedStack{}
}

func (s *SortedStack) Push(val int) {
	s.stack = append(s.stack, val)
	sort.Ints(s.stack)
}

func (s *SortedStack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.stack = s.stack[1:]
}

func (s *SortedStack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	return s.stack[0]
}

func (s *SortedStack) IsEmpty() bool {
	return len(s.stack) == 0
}
