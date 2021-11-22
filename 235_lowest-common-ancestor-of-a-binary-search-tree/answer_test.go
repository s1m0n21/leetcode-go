/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/30 5:12 下午
 */

package _lowest_common_ancestor_of_a_binary_search_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root, p, q *utils.TreeNode
	}

	root := utils.MakeBinaryTree(6,2,8,0,4,7,9,nil,nil,3,5)

	tests := []struct{
		input input
		expect *utils.TreeNode
	}{
		{input{root, root.Get(2), root.Get(8)}, root.Get(6)},
		{input{root, root.Get(2), root.Get(4)}, root.Get(2)},
	}

	for _, test := range tests {
		if actual := lowestCommonAncestor(test.input.root, test.input.p, test.input.q); actual != test.expect {
			t.Errorf("expect = %+v, actual = %+v", test.expect, actual)
		}
	}
}
