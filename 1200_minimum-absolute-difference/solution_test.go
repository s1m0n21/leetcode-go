/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/4 09:35
 */

package _minimum_absolute_difference

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, minimumAbsDifference)

	test.SetAndRun([]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}})
	test.SetAndRun([]int{1, 3, 6, 10, 15}, [][]int{{1, 3}})
	test.SetAndRun([]int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}})
	test.SetAndRun([]int{1, 4, 5, 6}, [][]int{{4, 5}, {5, 6}})
	test.SetAndRun([]int{1, 4}, [][]int{{1, 4}})
}
