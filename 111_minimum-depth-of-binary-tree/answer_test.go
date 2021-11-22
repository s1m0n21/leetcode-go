/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/26 3:35 下午
 */

package _minimum_depth_of_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeBinaryTree(3,9,20,nil,nil,15,7)

	t.Logf("answer = %d", minDepth(root))
}
