/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/21 17:55
 */

package _ugly_number

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, isUgly)

	testCase.SetAndRun(6, true)
	testCase.SetAndRun(8, true)
	testCase.SetAndRun(14, false)
	testCase.SetAndRun(1, true)
}
