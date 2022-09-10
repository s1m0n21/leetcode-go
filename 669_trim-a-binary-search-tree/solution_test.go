/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/9/10 12:52
 */

package _trim_a_binary_search_tree

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		root      *utils.TreeNode
		low, high int
	}

	tests := utils.NewTestCase(t, func(i input) *utils.TreeNode {
		return trimBST(i.root, i.low, i.high)
	})
	tests.ConfigSetCheckFunc(utils.SameTree)

	tests.SetAndRun(input{utils.MakeBinaryTree(1, 0, 2), 1, 2}, utils.MakeBinaryTree(1, nil, 2))
	tests.SetAndRun(input{utils.MakeBinaryTree(3, 0, 4, nil, 2, nil, nil, 1), 1, 3}, utils.MakeBinaryTree(3, 2, nil, 1))
}
