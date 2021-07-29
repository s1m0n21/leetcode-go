/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/29 11:09 下午
 */

package _path_sum_ii

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		root   *utils.TreeNode
		target int
	}

	tests := []struct {
		input  input
		expect [][]int
	}{
		{input{utils.MakeTreeNode(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1), 22}, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{input{utils.MakeTreeNode(1, 2, 3), 5}, [][]int{}},
		{input{utils.MakeTreeNode(1, 2), 0}, [][]int{}},
		{input{utils.MakeTreeNode(-2, nil, -3), -5}, [][]int{{-2,-3}}},
	}

	for _, test := range tests {
		if actual := pathSum(test.input.root, test.input.target); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("expect = %+v, actual = %+v", test.expect, actual)
		}
	}
}
