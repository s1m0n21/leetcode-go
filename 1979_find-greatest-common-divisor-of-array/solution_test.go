/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/16 10:20
 */

package _find_greatest_common_divisor_of_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, findGCD)

	testCase.SetAndRun([]int{2, 5, 6, 9, 10}, 2)
	testCase.SetAndRun([]int{7, 5, 6, 8, 3}, 1)
	testCase.SetAndRun([]int{3, 3}, 3)
}
