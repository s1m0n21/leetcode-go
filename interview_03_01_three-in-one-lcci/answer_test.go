/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/23 2:31 下午
 */

package interview_03_01_three_in_one_lcci

import "testing"

func TestAnswer(t *testing.T) {
	stacks := Constructor(3)

	stacks.Push(0, 1)
	stacks.Push(0, 2)
	stacks.Push(0, 3)
	stacks.Push(0, 4)
	t.Logf("%d", stacks.Pop(0))
	stacks.Push(0, 4)
	t.Logf("%d", stacks.Peek(0))
	t.Logf("%d", stacks.Peek(1))
	t.Logf("%t", stacks.IsEmpty(1))
	stacks.Push(1, 1)
	stacks.Push(1, 2)
	stacks.Push(1, 3)
	stacks.Push(2, 4)
	t.Logf("%d", stacks.Peek(1))
	t.Logf("%t", stacks.IsEmpty(1))
}
