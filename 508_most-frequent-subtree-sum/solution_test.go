/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/21 09:29
 */

package _most_frequent_subtree_sum

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, findFrequentTreeSum)
	test.ConfigSetOutputInAnyOrder()

	test.SetAndRun(utils.MakeBinaryTree(5, 2, -3), []int{2, -3, 4})
	test.SetAndRun(utils.MakeBinaryTree(5, 2, -5), []int{2})
}
