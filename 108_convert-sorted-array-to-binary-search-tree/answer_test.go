/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/1 2:32 上午
 */

package _convert_sorted_array_to_binary_search_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect *utils.TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, utils.MakeTreeNode(0, -3, 9, -10, nil, 5)},
		{[]int{1, 3}, utils.MakeTreeNode(3, 1)},
	}

	for _, test := range tests {

		if actual := sortedArrayToBST(test.input); !utils.SameTree(test.expect, actual) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, actual)
		}
	}
}
