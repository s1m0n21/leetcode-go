/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/30 4:13 下午
 */

package _path_sum_iii

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root   *utils.TreeNode
		target int
	}

	tests := []struct {
		input input
		expect int
	}{
		{input{utils.MakeTreeNode(10,5,-3,3,2,nil,11,3,-2,nil,1), 8}, 3},
		{input{utils.MakeTreeNode(5,4,8,11,nil,13,4,7,2,nil,nil,5,1), 22}, 3},
		{input{utils.MakeTreeNode(), 1}, 0},
	}

	for _, test := range tests {
		if actual := pathSum(test.input.root, test.input.target); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
