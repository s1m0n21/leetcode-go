/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/23 14:05
 */

package _n_ary_tree_preorder_traversal

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, preorder)

	testCase.SetAndRun(utils.MakeNaryTree(1, nil, 3, 2, 4, nil, 5, 6), []int{1, 3, 5, 6, 2, 4})
	testCase.SetAndRun(utils.MakeNaryTree(1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14), []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10})
}
