/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/7/25 09:52
 */

package _complete_binary_tree_inserter

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	cbt := Constructor(utils.MakeBinaryTree(1, 2, 3, 4, 5, 6))
	tests := utils.NewTestCase(t, cbt.Insert)

	tests.SetAndRun(7, 3)
	tests.SetAndRun(8, 4)

	//cbt.Get_root().Graph("cbt_after_insertion")
}
