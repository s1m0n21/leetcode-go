/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/16 02:37
 */

package interview_04_06_successor_lcci

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	testCase := utils.NewTestCase(t, func(i [2]*utils.TreeNode) *utils.TreeNode {
		return inorderSuccessor(i[0], i[1])
	})

	t1 := utils.MakeBinaryTree(2, 1, 3)
	testCase.SetAndRun([2]*utils.TreeNode{t1, t1.Get(1)}, t1.Get(2))

	t2 := utils.MakeBinaryTree(5, 3, 6, 2, 4, nil, nil, 1)
	testCase.SetAndRun([2]*utils.TreeNode{t2, t2.Get(6)}, nil)
}
