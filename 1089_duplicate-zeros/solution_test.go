/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/17 09:29
 */

package _duplicate_zeros

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, func(i []int) []int {
		duplicateZeros(i)
		return i
	})

	test.SetAndRun([]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4})
	test.SetAndRun([]int{1, 2, 3}, []int{1, 2, 3})
	test.SetAndRun([]int{0, 1}, []int{0, 0})
}
