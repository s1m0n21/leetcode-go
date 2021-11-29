/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/29 9:59 上午
 */

package _design_circular_queue

import "testing"

func TestAnswer(t *testing.T) {
	queue := Constructor(3)

	t.Logf("%t; true", queue.EnQueue(1))
	t.Logf("%t; true", queue.EnQueue(2))
	t.Logf("%t; true", queue.EnQueue(3))
	t.Logf("%t; false", queue.EnQueue(4))
	t.Logf("%d; 3", queue.Rear())
	t.Logf("%t; true", queue.IsFull())
	t.Logf("%t; true", queue.DeQueue())
	t.Logf("%t; true", queue.EnQueue(4))
	t.Logf("%d; 4", queue.Rear())
	t.Logf("%d; 2", queue.Front())
}
