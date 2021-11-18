/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/18 1:56 下午
 */

package _binary_tree_tilt

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.TreeNode
		expect int
	}{
		{utils.MakeTreeNode(1, 2, 3), 1},
		{utils.MakeTreeNode(4, 2, 9, 3, 5, nil, 7), 15},
		{utils.MakeTreeNode(21, 7, 14, 1, 1, 2, 2, 3, 3), 9},
		{utils.MakeTreeNode(), 0},
	}

	for i, test := range tests {
		if actual := findTilt(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
