/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/31 14:06
 */

package _self_dividing_numbers

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		left, right int
	}
	testCase := utils.NewTestCase(t, func(i input) []int {
		return selfDividingNumbers(i.left, i.right)
	})

	testCase.SetAndRun(input{1, 22}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22})
	testCase.SetAndRun(input{47, 85}, []int{48, 55, 66, 77})
}
