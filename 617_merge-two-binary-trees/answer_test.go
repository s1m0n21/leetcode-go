/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/24 09:37
 */

package _merge_two_binary_trees

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, func(roots [2]*utils.TreeNode) []int {
		return mergeTrees(roots[0], roots[1]).Preorder()
	})

	testCase.SetAndRun([2]*utils.TreeNode{utils.MakeBinaryTree(1, 3, 2, 5), utils.MakeBinaryTree(2, 1, 3, nil, 4, nil, 7)}, utils.MakeBinaryTree(3, 4, 5, 5, 4, nil, 7).Preorder())
	testCase.SetAndRun([2]*utils.TreeNode{utils.MakeBinaryTree(1), utils.MakeBinaryTree(1, 2)}, utils.MakeBinaryTree(2, 2).Preorder())
}
