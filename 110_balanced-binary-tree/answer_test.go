/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/27 4:26 下午
 */

package _balanced_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(1,2,2,3,3,nil,nil,4,4)

	t.Logf("answer = %t", isBalanced(root))
}
