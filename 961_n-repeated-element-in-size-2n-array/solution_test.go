/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/21 09:32
 */

package _n_repeated_element_in_size_2n_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, repeatedNTimes)

	testCase.SetAndRun([]int{1, 2, 3, 3}, 3)
	testCase.SetAndRun([]int{2, 1, 2, 5, 3, 2}, 2)
	testCase.SetAndRun([]int{5, 1, 5, 2, 5, 3, 5, 4}, 5)
	testCase.SetAndRun([]int{1, 1, 2, 3}, 1)
}
