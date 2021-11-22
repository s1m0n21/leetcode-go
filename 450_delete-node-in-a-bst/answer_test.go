/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/31 4:26 上午
 */

package _delete_node_in_a_bst

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root *utils.TreeNode
		key  int
	}

	tests := []struct {
		input  input
		expect *utils.TreeNode
	}{
		{input{utils.MakeBinaryTree(5, 3, 6, 2, 4, nil, 7), 3}, utils.MakeBinaryTree(5, 4, 6, 2, nil, nil, 7)},
	}

	for _, test := range tests {
		if actual := deleteNode(test.input.root, test.input.key); !utils.SameTree(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
