/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/5 14:01
 */

package _squares_of_a_sorted_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, func(i []int) []int {
		c := append([]int(nil), i...)
		return sortedSquares(c)
	})

	testCase.SetAndRun([]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100})
	testCase.SetAndRun([]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121})
}
