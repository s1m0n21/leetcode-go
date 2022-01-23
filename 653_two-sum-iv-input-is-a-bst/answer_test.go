/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/23 12:07 PM
 */

package _two_sum_iv_input_is_a_bst

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root *utils.TreeNode
		k    int
	}

	tests := []struct {
		input  input
		expect bool
	}{
		{input{utils.MakeBinaryTree(5, 3, 6, 2, 4, nil, 7), 9}, true},
		{input{utils.MakeBinaryTree(5, 3, 6, 2, 4, nil, 7), 28}, false},
	}

	for i, test := range tests {
		if actual := findTarget(test.input.root, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
