/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/22 13:51
 */

package _average_of_levels_in_binary_tree

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, averageOfLevels)

	testCase.SetAndRun(utils.MakeBinaryTree(3, 9, 20, nil, nil, 15, 7), []float64{3, 14.5, 11})
	testCase.SetAndRun(utils.MakeBinaryTree(3, 9, 20, 15, 7), []float64{3, 14.5, 11})
	testCase.SetAndRun(utils.MakeBinaryTree(1), []float64{1})
}
