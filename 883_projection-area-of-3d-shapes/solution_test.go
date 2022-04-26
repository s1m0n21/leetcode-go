/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/26 13:13
 */

package _projection_area_of_3d_shapes

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, projectionArea)

	testCase.SetAndRun([][]int{{1, 2}, {3, 4}}, 17)
	testCase.SetAndRun([][]int{{2}}, 5)
	testCase.SetAndRun([][]int{{1, 0}, {0, 2}}, 8)
}
