/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/22 11:22
 */

package _construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		preorder, inorder []int
	}

	testCase := utils.NewTestCase(t, func(input input) []int {
		return buildTree(input.preorder, input.inorder).Preorder()
	})

	testCase.SetAndRun(input{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}}, utils.MakeBinaryTree(3, 9, 20, nil, nil, 15, 7).Preorder())
	testCase.SetAndRun(input{[]int{-1}, []int{-1}}, utils.MakeBinaryTree(-1).Preorder())
	testCase.SetAndRun(input{[]int{}, []int{}}, nil)
}
