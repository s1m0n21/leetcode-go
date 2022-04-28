/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/27 20:11
 */

package offer_18_shan_chu_lian_biao_de_jie_dian_lcof

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		head *utils.ListNode
		val  int
	}

	testCase := utils.NewTestCase(t, func(i input) *utils.ListNode {
		return deleteNode(i.head, i.val)
	}, utils.OptionCheckFunc(utils.SameList))

	testCase.SetAndRun(input{utils.MakeListNode(1, 2, 3, 4), 3}, utils.MakeListNode(1, 2, 4))
	testCase.SetAndRun(input{utils.MakeListNode(1, 2, 3, 4), 1}, utils.MakeListNode(2, 3, 4))
}
