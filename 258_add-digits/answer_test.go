/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/22 11:51
 */

package _add_digits

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, addDigits)

	testCase.SetAndRun(38, 2)
	testCase.SetAndRun(0, 0)
	testCase.SetAndRun(99, 9)
	testCase.SetAndRun(1<<31-1, 1)
}
