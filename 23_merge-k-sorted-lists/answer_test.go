/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/25 11:57 下午
 */

package _merge_k_sorted_lists

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	lists := []*utils.ListNode{
		utils.MakeListNode(1,2,3,4),
		utils.MakeListNode(5,6,7,8),
		utils.MakeListNode(9,10,11,12),
		utils.MakeListNode(13,14,15,16),
	}
	res := mergeKLists(lists)

	t.Logf("answer = %s", res)
}
