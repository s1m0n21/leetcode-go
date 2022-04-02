/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/23 4:36 下午
 */

package _binary_tree_level_order_traversal

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, levelOrder)

	testCase.SetAndRun(utils.MakeBinaryTree(3, 9, 20, nil, nil, 15, 7), [][]int{{3}, {9, 20}, {15, 7}})
	testCase.SetAndRun(utils.MakeBinaryTree(1), [][]int{{1}})
	testCase.SetAndRun(utils.MakeBinaryTree(), nil)
}
