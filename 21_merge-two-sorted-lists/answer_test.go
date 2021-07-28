/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/19 4:22 下午
 */

package _merge_two_sorted_lists

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	l1 := utils.MakeListNode(1,2,4)
	l2 := utils.MakeListNode(1,3,4)
	res := mergeTwoLists(l1, l2)

	t.Logf("answer = %s", res)
}
