/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/3 12:30 上午
 */

package _min_stack

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	stack := Constructor()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	utils.Log.Debugf("min = %d", stack.GetMin())

	stack.Pop()

	utils.Log.Debugf("top = %d", stack.Top())
	utils.Log.Debugf("min = %d", stack.GetMin())
}
