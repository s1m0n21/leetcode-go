/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/25 17:06
 */

package _find_all_numbers_disappeared_in_an_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	testCase := utils.NewTestCase(t, func(n []int) []int {
		c := append([]int(nil), n...)
		return findDisappearedNumbers(c)
	})

	testCase.SetAndRun([]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6})
	testCase.SetAndRun([]int{1, 1}, []int{2})
}
