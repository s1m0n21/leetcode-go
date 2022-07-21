/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/21 09:48
 */

package _binary_tree_pruning

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	tests := utils.NewTestCase(t, pruneTree)
	tests.ConfigSetCheckFunc(utils.SameTree)

	tests.SetAndRun(utils.MakeBinaryTree(1, nil, 0, 0, 1), utils.MakeBinaryTree(1, nil, 0, nil, 1))
	tests.SetAndRun(utils.MakeBinaryTree(1, 0, 1, 0, 0, 0, 1), utils.MakeBinaryTree(1, nil, 1, nil, 1))
	tests.SetAndRun(utils.MakeBinaryTree(1, 1, 0, 1, 1, 0, 1, 0), utils.MakeBinaryTree(1, 1, 0, 1, 1, nil, 1))
	tests.SetAndRun(utils.MakeBinaryTree(0, nil, 0, 0, 0), nil)
}
