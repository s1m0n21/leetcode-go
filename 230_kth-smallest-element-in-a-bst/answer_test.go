/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/2 4:22 上午
 */

package _kth_smallest_element_in_a_bst

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
		expect int
	}{
		{input{utils.MakeTreeNode(3, 1, 4, nil, 2), 1}, 1},
		{input{utils.MakeTreeNode(5, 3, 6, 2, 4, nil, nil, 1), 3}, 3},
	}

	for _, test := range tests {
		if actual := kthSmallest(test.input.root, test.input.k); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
