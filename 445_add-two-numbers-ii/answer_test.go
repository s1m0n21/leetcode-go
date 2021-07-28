/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/19 12:42 上午
 */

package _add_two_numbers_ii

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	l1 := utils.MakeListNode(5)
	l2 := utils.MakeListNode(5)
	res := addTwoNumbers(l1, l2)

	t.Logf("answer = %s", res)
}
