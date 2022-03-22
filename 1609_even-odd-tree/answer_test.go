/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/20 4:13 下午
 */

package _even_odd_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, isEvenOddTree)

	testCase.SetAndRun(utils.MakeBinaryTree(1, 10, 4, 3, nil, 7, 9, 12, 8, 6, nil, nil, 2), true)
	testCase.SetAndRun(utils.MakeBinaryTree(5, 4, 2, 3, 3, 7), false)
	testCase.SetAndRun(utils.MakeBinaryTree(5, 9, 1, 3, 5, 7), false)
	testCase.SetAndRun(utils.MakeBinaryTree(11, 8, 6, 1, 3, 9, 11, 30, 20, 18, 16, 12, 10, 4, 2, 17), true)
	testCase.SetAndRun(utils.MakeBinaryTree(1), true)
}
