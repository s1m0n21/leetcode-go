/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/22 17:00
 */

package _fibonacci_number

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, fib)

	testCase.SetAndRun(2, 1)
	testCase.SetAndRun(3, 2)
	testCase.SetAndRun(4, 3)
}
