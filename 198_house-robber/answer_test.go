/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/2 17:18
 */

package _house_robber

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, rob)

	testCase.SetAndRun([]int{1, 2, 3, 1}, 4)
	testCase.SetAndRun([]int{2, 7, 9, 3, 1}, 12)
	testCase.SetAndRun([]int{9}, 9)
	testCase.SetAndRun([]int{2, 1, 1, 2}, 4)
}
