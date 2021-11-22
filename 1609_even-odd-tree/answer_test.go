/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/20 4:13 下午
 */

package _even_odd_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.TreeNode
		expect bool
	}{
		{utils.MakeBinaryTree(1, 10, 4, 3, nil, 7, 9, 12, 8, 6, nil, nil, 2), true},
		{utils.MakeBinaryTree(5, 4, 2, 3, 3, 7), false},
		{utils.MakeBinaryTree(5, 9, 1, 3, 5, 7), false},
		{utils.MakeBinaryTree(11, 8, 6, 1, 3, 9, 11, 30, 20, 18, 16, 12, 10, 4, 2, 17), true},
		{utils.MakeBinaryTree(1), true},
	}

	for i, test := range tests {
		if actual := isEvenOddTree(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %t, actual = %t", i, test.input.Preorder(), test.expect, actual)
		}
	}
}
