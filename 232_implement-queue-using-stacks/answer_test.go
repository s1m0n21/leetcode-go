/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/11 2:21 上午
 */

package _implement_queue_using_stacks

import "testing"

func TestAnswer(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)

	t.Logf("peek: %d", q.Peek())
	t.Logf("pop: %d", q.Pop())
	t.Logf("empty: %t", q.Empty())
	t.Logf("pop: %d", q.Pop())
	t.Logf("empty: %t", q.Empty())
}
