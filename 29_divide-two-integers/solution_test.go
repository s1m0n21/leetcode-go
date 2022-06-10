/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/10 09:45
 */

package _divide_two_integers

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, func(i [2]int) int {
		return divide(i[0], i[1])
	})

	test.SetAndRun([2]int{10, 5}, 10/5)
	test.SetAndRun([2]int{100, 2}, 100/2)
	test.SetAndRun([2]int{81, -1}, 81/-1)
	test.SetAndRun([2]int{(1 << 31) - 1, 1}, ((1<<31)-1)/1)
	test.SetAndRun([2]int{-2147483648, -1}, (1<<31)-1)
}
