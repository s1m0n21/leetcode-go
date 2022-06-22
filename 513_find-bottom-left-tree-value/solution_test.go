/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/22 10:15
 */

package _find_bottom_left_tree_value

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, findBottomLeftValue)

	test.SetAndRun(utils.MakeBinaryTree(2, 1, 3), 1)
	test.SetAndRun(utils.MakeBinaryTree(1, 2, 3, 4, nil, 5, 6, nil, nil, 7), 7)
	test.SetAndRun(utils.MakeBinaryTree(1), 1)
}
