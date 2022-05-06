/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/3 13:49
 */

package _minimum_falling_path_sum

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, minFallingPathSum)

	testCase.SetAndRun([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13)
	testCase.SetAndRun([][]int{{-19, 57}, {-40, -5}}, -59)
	testCase.SetAndRun([][]int{{1}}, 1)
	testCase.SetAndRun([][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}}, -36)
}
