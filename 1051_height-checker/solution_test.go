/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/13 01:02
 */

package _height_checker

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, heightChecker)

	test.SetAndRun([]int{1, 1, 4, 2, 1, 3}, 3)
	test.SetAndRun([]int{5, 1, 2, 3, 4}, 5)
	test.SetAndRun([]int{1, 2, 3, 4, 5}, 0)
}
