/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/22 11:15 上午
 */

package _maximum_depth_of_n_ary_tre

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  *utils.NTreeNode
		expect int
	}{
		{utils.MakeNaryTree(1, nil, 3, 2, 4, nil, 5, 6), 3},
		{utils.MakeNaryTree(1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14), 5},
		{utils.MakeNaryTree(), 0},
		{utils.MakeNaryTree(1), 1},
	}

	for i, test := range tests {
		if actual := maxDepth(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input.Preorder(), test.expect, actual)
		}
	}
}
