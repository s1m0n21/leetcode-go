/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/26 9:46 上午
 */

package _search_in_a_binary_search_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root *utils.TreeNode
		val  int
	}

	tree := utils.MakeBinaryTree(4, 2, 7, 1, 3)

	tests := []struct {
		input  input
		expect *utils.TreeNode
	}{
		{input{tree, 2}, tree.Get(2)},
		{input{tree, 5}, nil},
	}

	for i, test := range tests {
		if actual := searchBST(test.input.root, test.input.val); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect.Preorder(), actual.Preorder())
		}
	}
}
