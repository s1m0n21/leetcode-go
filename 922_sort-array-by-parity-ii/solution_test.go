/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/16 11:05
 */

package _sort_array_by_parity_ii

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, sortArrayByParityII)

	testCase.SetAndRun([]int{4, 2, 5, 7}, []int{4, 5, 2, 7})
	testCase.SetAndRun([]int{2, 3}, []int{2, 3})
	testCase.SetAndRun([]int{3, 2}, []int{2, 3})
}
