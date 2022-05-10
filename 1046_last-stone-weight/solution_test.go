/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/10 10:07
 */

package _last_stone_weight

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, lastStoneWeight)

	testCase.SetAndRun([]int{2, 7, 4, 1, 8, 1}, 1)
	testCase.SetAndRun([]int{2}, 2)
	testCase.SetAndRun([]int{2, 2}, 0)
	testCase.SetAndRun([]int{2, 2, 3}, 1)
}
