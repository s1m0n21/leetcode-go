/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/11 17:17
 */

package _find_pivot_index

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, pivotIndex)

	tests.SetAndRun([]int{1, 7, 3, 6, 5, 6}, 3)
	tests.SetAndRun([]int{1, 2, 3}, -1)
	tests.SetAndRun([]int{2, 1, -1}, 0)
	tests.SetAndRun([]int{-1, 1, 2}, 2)
	tests.SetAndRun([]int{0}, 0)
}
