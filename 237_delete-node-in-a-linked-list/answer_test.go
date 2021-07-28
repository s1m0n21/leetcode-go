/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/20 8:19 下午
 */

package _delete_node_in_a_linked_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1,2,3,4,5)
	del := head
	for del.Next != nil && del.Next.Next != nil {
		del = del.Next
	}
	deleteNode(del)

	t.Logf("answer = %s", head)
}
