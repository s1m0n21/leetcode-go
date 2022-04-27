/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/4/27 20:00
 */

package offer_06_cong_wei_dao_tou_da_yin_lian_biao_lcof

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, reversePrint)

	testCase.SetAndRun(utils.MakeListNode(1, 3, 2), []int{2, 3, 1})
	testCase.SetAndRun(utils.MakeListNode(), nil)
}
