/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/9 5:15 下午
 */

package interview_17_12_binode_lcci

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeBinaryTree(4, 2, 5, 1, 3, nil, 6, 0)
	head := convertBiNode(root)
	t.Logf("%+v", head.Preorder())
}
