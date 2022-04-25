/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/25 23:08
 */

package _98_random_pick_index

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	s := Constructor([]int{1, 2, 3, 3, 3})
	testCase := utils.NewTestCase(t, s.Pick)

	testCase.SetAndRun(3, 4)
	testCase.SetAndRun(1, 0)
}
