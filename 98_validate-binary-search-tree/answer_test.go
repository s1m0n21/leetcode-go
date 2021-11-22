/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/30 5:32 下午
 */

package _validate_binary_search_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct{
		input *utils.TreeNode
		expect bool
	}{
		{utils.MakeBinaryTree(2,1,3), true},
		{utils.MakeBinaryTree(5,1,4,nil,nil,3,6), false},
		{utils.MakeBinaryTree(2,2,2), false},
		{utils.MakeBinaryTree(5,4,6,nil,nil,3,7), false},
	}

	for _, test := range tests {
		if actual := isValidBST(test.input); actual != test.expect {
			t.Errorf("input = %+v, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
