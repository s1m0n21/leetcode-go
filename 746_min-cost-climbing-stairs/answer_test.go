/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/4 17:07
 */

package _min_cost_climbing_stairs

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, minCostClimbingStairs)

	testCase.SetAndRun([]int{10, 15, 20}, 15)
	testCase.SetAndRun([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6)
	testCase.SetAndRun([]int{1, 1}, 1)
}
