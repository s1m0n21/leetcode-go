/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/28 17:43
 */

package _n_queens_ii

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, totalNQueens)

	testCase.SetAndRun(4, 2)
	testCase.SetAndRun(1, 1)
}
