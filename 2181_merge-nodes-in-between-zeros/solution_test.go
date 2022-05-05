/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/5 10:09
 */

package _merge_nodes_in_between_zeros

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, mergeNodes, utils.OptionCheckFunc(utils.SameList))

	testCase.SetAndRun(utils.MakeListNode(0, 3, 1, 0, 4, 5, 2, 0), utils.MakeListNode(4, 11))
	testCase.SetAndRun(utils.MakeListNode(0, 1, 0, 3, 0, 2, 2, 0), utils.MakeListNode(1, 3, 4))
}
