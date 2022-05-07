/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/8 01:29
 */

package _find_all_duplicates_in_an_array

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, findDuplicates)

	testCase.SetAndRun([]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3})
	testCase.SetAndRun([]int{1, 1, 2}, []int{1})
	testCase.SetAndRun([]int{1}, nil)
}
