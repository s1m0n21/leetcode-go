/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/27 11:21 上午
 */

package _second_minimum_node_in_a_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(2,2,5,nil,nil,5,5)

	t.Logf("answer = %d", findSecondMinimumValue(root))
}
