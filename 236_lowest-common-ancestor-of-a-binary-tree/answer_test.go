/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/7 11:17 AM
 */

package _lowest_common_ancestor_of_a_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root, p, q *utils.TreeNode
	}

	t1 := utils.MakeBinaryTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
	t2 := utils.MakeBinaryTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
	t3 := utils.MakeBinaryTree(1, 2)

	tests := []struct {
		input  input
		expect *utils.TreeNode
	}{
		{input{t1, t1.Get(5), t1.Get(1)}, t1.Get(3)},
		{input{t2, t2.Get(5), t2.Get(4)}, t2.Get(5)},
		{input{t3, t3.Get(1), t3.Get(2)}, t3.Get(1)},
	}

	for i, test := range tests {
		if actual := lowestCommonAncestor(test.input.root, test.input.p, test.input.q); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
