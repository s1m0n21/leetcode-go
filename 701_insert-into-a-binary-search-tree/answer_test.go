/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/23 12:24 PM
 */

package _insert_into_a_binary_search_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root *utils.TreeNode
		val  int
	}

	tests := []struct {
		input  input
		expect *utils.TreeNode
	}{
		{input{utils.MakeBinaryTree(4, 2, 7, 1, 3), 5}, utils.MakeBinaryTree(4, 2, 7, 1, 3, 5)},
		{input{utils.MakeBinaryTree(40, 20, 60, 10, 30, 50, 70), 25}, utils.MakeBinaryTree(40, 20, 60, 10, 30, 50, 70, nil, nil, 25)},
		{input{utils.MakeBinaryTree(4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil), 5}, utils.MakeBinaryTree(4, 2, 7, 1, 3, 5, nil, nil, nil, nil, nil)},
	}

	for i, test := range tests {
		if actual := insertIntoBST(test.input.root, test.input.val); !utils.SameTree(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
