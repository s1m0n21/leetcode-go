/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/4 1:23 上午
 */

package _convert_sorted_list_to_binary_search_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.ListNode
		expect *utils.TreeNode
	}{
		{utils.MakeListNode(-10, -3, 0, 5, 9), utils.MakeTreeNode(0, -3, 9, -10, nil, 5)},
		{utils.MakeListNode(1, 2, 3, 4, 5, 6, 7), utils.MakeTreeNode(4, 2, 6, 1, 3, 5, 7)},
	}

	for _, test := range tests {
		if actual := sortedListToBST(test.input); !utils.SameTree(actual, test.expect) {
			t.Errorf("input = %s, expect = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
