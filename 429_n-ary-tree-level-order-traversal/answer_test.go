/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/5 1:18 PM
 */

package _n_ary_tree_level_order_traversal

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.NTreeNode
		expect [][]int
	}{
		{utils.MakeNaryTree(1, nil, 3, 2, 4, nil, 5, 6), [][]int{{1}, {3, 2, 4}, {5, 6}}},
		{utils.MakeNaryTree(1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14), [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}}},
	}

	for i, test := range tests {
		if actual := levelOrder(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input.Preorder(), test.expect, actual)
		}
	}
}
