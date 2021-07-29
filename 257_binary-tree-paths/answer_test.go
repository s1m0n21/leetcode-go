/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/29 10:08 下午
 */

package _binary_tree_paths

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.TreeNode
		expect []string
	}{
		{utils.MakeTreeNode(1, 2, 3, nil, 5), []string{"1->2->5", "1->3"}},
		{utils.MakeTreeNode(1, 2, 3, 4, 5), []string{"1->2->4", "1->2->5", "1->3"}},
	}

	for _, test := range tests {
		if actual := binaryTreePaths(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input.Preorder(), test.expect, actual)
		}
	}
}
