/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/31 17:50
 */

package _diagonal_traverse_ii

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, findDiagonalOrder)

	testCase.SetAndRun([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 4, 2, 7, 5, 3, 8, 6, 9})
	testCase.SetAndRun([][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}, []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16})
	testCase.SetAndRun([][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}, []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11})
	testCase.SetAndRun([][]int{{1, 2, 3, 4, 5, 6}}, []int{1, 2, 3, 4, 5, 6})
}
