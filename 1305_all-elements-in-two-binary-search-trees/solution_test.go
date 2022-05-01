/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/1 13:36
 */

package _all_elements_in_two_binary_search_trees

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	type input struct {
		a, b *utils.TreeNode
	}

	testCase := utils.NewTestCase(t, func(i input) []int {
		return getAllElements(i.a, i.b)
	})

	testCase.SetAndRun(input{utils.MakeBinaryTree(2, 1, 4), utils.MakeBinaryTree(1, 0, 3)}, []int{0, 1, 1, 2, 3, 4})
	testCase.SetAndRun(input{utils.MakeBinaryTree(1, nil, 8), utils.MakeBinaryTree(8, 1)}, []int{1, 1, 8, 8})
}
