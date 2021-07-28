/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 9:04 下午
 */

package _add_two_numbers

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	l1 := utils.MakeListNode(2, 4, 3)
	l2 := utils.MakeListNode(5, 6, 4)

	res := addTwoNumbers(l1, l2)
	for res != nil {
		t.Logf("answer = %d", res.Val)
		res = res.Next
	}
}
