/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/30 10:09
 */

package _sum_of_root_to_leaf_binary_numbers

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, sumRootToLeaf)

	test.SetAndRun(utils.MakeBinaryTree(1, 0, 1, 0, 1, 0, 1), 22)
	test.SetAndRun(utils.MakeBinaryTree(0), 0)
}
