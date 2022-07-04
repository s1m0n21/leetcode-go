/**
 * @Author         : s1m0n21
 * @Description    : TestSolution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/4 09:56
 */

package _next_greater_element_iii

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, nextGreaterElement)

	test.SetAndRun(12, 21)
	test.SetAndRun(21, -1)
	test.SetAndRun(123, 321)
	test.SetAndRun(1<<31-1, -1)
}
