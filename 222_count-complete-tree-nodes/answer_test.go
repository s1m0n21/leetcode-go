/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/26 6:06 下午
 */

package _count_complete_tree_nodes

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(1,2,3,4,5,6)

	t.Logf("answer = %d", countNodes(root))
}
