/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/4 17:29
 */

package interview_03_05_sort_of_stacks_lcci

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	stack := Constructor()
	testCase := utils.NewTestCase(t, func(i int) int {
		stack.Push(i)
		return stack.Peek()
	})

	testCase.SetAndRun(1, 1)
	testCase.SetAndRun(2, 1)
	testCase.SetAndRun(3, 1)
	stack.Pop()
	testCase.SetAndRun(4, 2)
	testCase.SetAndRun(0, 0)
}
