/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/5/24 09:22
 */

package _univalued_binary_tree

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
	test := utils.NewTestCase(t, isUnivalTree)

	test.SetAndRun(utils.MakeBinaryTree(1, 1, 1, 1, 1, nil, 1), true)
	test.SetAndRun(utils.MakeBinaryTree(2, 2, 2, 5, 2), false)
}
