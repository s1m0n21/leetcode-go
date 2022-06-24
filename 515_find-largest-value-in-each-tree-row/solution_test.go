/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/24 13:23
 */

package _find_largest_value_in_each_tree_row

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, largestValues)

	test.SetAndRun(utils.MakeBinaryTree(1, 3, 2, 5, 3, nil, 9), []int{1, 3, 9})
	test.SetAndRun(utils.MakeBinaryTree(1, 2, 3), []int{1, 3})
	test.SetAndRun(utils.MakeBinaryTree(-1, -2, -3, nil, nil, 0), []int{-1, -2, 0})
	test.SetAndRun(utils.MakeBinaryTree(), nil)
	test.SetAndRun(utils.MakeBinaryTree(1), []int{1})
}
