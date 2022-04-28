/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/28 16:16
 */

package _linked_list_components

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		head *utils.ListNode
		nums []int
	}

	testCase := utils.NewTestCase(t, func(i input) int {
		return numComponents(i.head, i.nums)
	})

	testCase.SetAndRun(input{utils.MakeListNode(0, 1, 2, 3), []int{0, 1, 3}}, 2)
	testCase.SetAndRun(input{utils.MakeListNode(0, 1, 2, 3, 4), []int{0, 3, 1, 4}}, 2)
	testCase.SetAndRun(input{utils.MakeListNode(0), []int{1}}, 0)
}
