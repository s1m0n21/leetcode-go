/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/16 10:11
 */

package _k_diff_pairs_in_an_array

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		nums []int
		k    int
	}

	test := utils.NewTestCase(t, func(i input) int {
		return findPairs(i.nums, i.k)
	})

	test.SetAndRun(input{[]int{3, 1, 4, 1, 5}, 2}, 2)
	test.SetAndRun(input{[]int{1, 2, 3, 4, 5}, 1}, 4)
	test.SetAndRun(input{[]int{1, 3, 1, 5, 4}, 0}, 1)
	test.SetAndRun(input{[]int{1, 1, 1}, 0}, 1)
	test.SetAndRun(input{[]int{1}, 1}, 0)
}
