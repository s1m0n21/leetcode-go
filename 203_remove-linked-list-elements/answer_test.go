/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/5 7:43 下午
 */

package _remove_linked_list_elements

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	ln := utils.MakeListNode(0, 1, 2, 3, 4, 5, 6)

	res := removeElements(ln, 3)
	for res != nil {
		t.Logf("answer = %d", res.Val)
		res = res.Next
	}
}
