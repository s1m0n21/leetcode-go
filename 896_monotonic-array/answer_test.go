/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/28 10:17
 */

package _monotonic_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, isMonotonic)

	testCase.SetAndRun([]int{1, 2, 2, 3}, true)
	testCase.SetAndRun([]int{6, 5, 4, 4}, true)
	testCase.SetAndRun([]int{1, 3, 2}, false)
	testCase.SetAndRun([]int{1}, true)
}
