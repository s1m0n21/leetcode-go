/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/8 11:06
 */

package _minimum_cost_to_move_chips_to_the_same_position

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, minCostToMoveChips)

	tests.SetAndRun([]int{1, 2, 3}, 1)
	tests.SetAndRun([]int{2, 2, 2, 3, 3}, 2)
	tests.SetAndRun([]int{1, 1000000000}, 1)
}
