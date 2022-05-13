/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/13 10:04
 */

package interview_01_05_one_away_lcci

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, func(i [2]string) bool {
		return oneEditAway(i[0], i[1])
	})

	testCase.SetAndRun([2]string{"pale", "ple"}, true)
	testCase.SetAndRun([2]string{"pales", "pal"}, false)
	testCase.SetAndRun([2]string{"pales", "paaas"}, false)
	testCase.SetAndRun([2]string{"a", "bbbbbbb"}, false)
	testCase.SetAndRun([2]string{"abbbb", "c"}, false)
	testCase.SetAndRun([2]string{"a", "a"}, true)
	testCase.SetAndRun([2]string{"ab", "a"}, true)
	testCase.SetAndRun([2]string{"", "a"}, true)
	testCase.SetAndRun([2]string{"", ""}, true)
	testCase.SetAndRun([2]string{"abcdefg", "abcdfeg"}, false)
	testCase.SetAndRun([2]string{"abcdefg", "abcdeg"}, true)
	testCase.SetAndRun([2]string{"mart", "karma"}, false)
	testCase.SetAndRun([2]string{"karn", "karma"}, false)
}
