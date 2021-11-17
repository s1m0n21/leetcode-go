/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/17 3:06 下午
 */

package _delete_nodes_and_return_forest

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	root := utils.MakeTreeNode(1, 2, 3, 4, 5, 6, 7)
	forest := delNodes(root, []int{3, 5}) // expect contains: [[1,2,4],[6],[7]]
	for _, root := range forest {
		t.Logf("%+v", root.Preorder())
	}
}
