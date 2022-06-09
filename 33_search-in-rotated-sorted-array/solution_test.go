/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/9 12:00
 */

package _search_in_rotated_sorted_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	test := utils.NewTestCase(t, func(i input) int {
		return search(i.nums, i.target)
	})

	test.SetAndRun(input{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4)
	test.SetAndRun(input{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1)
	test.SetAndRun(input{[]int{1}, 0}, -1)
	test.SetAndRun(input{[]int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6}, 6}, 9)
}
