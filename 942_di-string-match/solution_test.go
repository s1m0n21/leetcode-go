/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/9 09:44
 */

package _di_string_match

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, diStringMatch)

	testCase.SetAndRun("IDID", []int{0, 4, 1, 3, 2})
	testCase.SetAndRun("III", []int{0, 1, 2, 3})
	testCase.SetAndRun("DDI", []int{3, 2, 0, 1})
	testCase.SetAndRun("DDD", []int{3, 2, 1, 0})
	testCase.SetAndRun("IID", []int{0, 1, 3, 2})
}
