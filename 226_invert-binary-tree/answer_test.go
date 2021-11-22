/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/26 3:49 下午
 */

package _invert_binary_tree

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeBinaryTree(4,2,7,1,3,6,9)
	invertTree(root)
	root.Graph()
}
