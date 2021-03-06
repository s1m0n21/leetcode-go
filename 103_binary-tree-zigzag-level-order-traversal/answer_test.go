/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/23 5:10 下午
 */

package _binary_tree_zigzag_level_order_traversal

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeBinaryTree(3,9,nil,nil,20,15,nil,nil,7)

	t.Logf("answer = %+v", zigzagLevelOrder(root))
}
