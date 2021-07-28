/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/20 8:44 下午
 */

package _remove_nth_node_from_end_of_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	head := utils.MakeListNode(1,2,3,4,5)
	n := 2
	res := removeNthFromEnd(head, n)

	t.Logf("answer = %s", res)
}
