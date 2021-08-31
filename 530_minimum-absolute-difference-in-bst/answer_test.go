/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/31 4:28 下午
 */

package _minimum_absolute_difference_in_bst

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.TreeNode
		expect int
	}{
		{utils.MakeTreeNode(1, nil, 3, 2), 1},
	}

	for _, test := range tests {
		if actual := getMinimumDifference(test.input); actual != test.expect {
			t.Errorf("input = %+v, exepct = %d, actual = %d", test.input.Preorder(), test.expect, actual)
		}
	}
}
