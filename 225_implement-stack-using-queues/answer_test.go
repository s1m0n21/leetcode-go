/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/1 5:23 下午
 */

package _implement_stack_using_queues

import "testing"

func TestAnswer(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	t.Logf("pop = %d", stack.Pop())
	t.Logf("pop = %d", stack.Pop())
	t.Logf("pop = %d", stack.Pop())
	t.Logf("top = %d", stack.Top())
	t.Logf("empty = %t", stack.Empty())
	t.Logf("pop = %d", stack.Pop())
	t.Logf("pop = %d", stack.Pop())
	t.Logf("empty = %t", stack.Empty())
}
