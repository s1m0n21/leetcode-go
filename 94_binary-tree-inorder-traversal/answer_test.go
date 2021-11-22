/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/23 11:55 上午
 */

package _binary_tree_inorder_traversal

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeBinaryTree(1,nil,2,3)

	t.Logf("answer = %+v", inorderTraversal(root))
}
