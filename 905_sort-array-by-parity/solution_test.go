/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/28 11:21
 */

package _sort_array_by_parity

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, sortArrayByParity)

	testCase.SetAndRun([]int{3, 1, 2, 4}, []int{4, 2, 1, 3})
	testCase.SetAndRun([]int{0}, []int{0})
	testCase.SetAndRun([]int{2, 4, 3, 6}, []int{2, 4, 6, 3})
}
