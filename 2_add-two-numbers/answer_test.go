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
	type input struct {
		l1, l2 *utils.ListNode
	}

	tests := []struct {
		input  input
		expect *utils.ListNode
	}{
		{input{utils.MakeListNode(2, 4, 3), utils.MakeListNode(5, 6, 4)}, utils.MakeListNode(7, 0, 8)},
		{input{utils.MakeListNode(3, 2, 1), utils.MakeListNode(4, 5, 6)}, utils.MakeListNode(7, 7, 7)},
		{input{utils.MakeListNode(9, 9, 9), utils.MakeListNode(2)}, utils.MakeListNode(1, 0, 0, 1)},
		{input{utils.MakeListNode(), utils.MakeListNode()}, utils.MakeListNode()},
	}

	for _, test := range tests {
		if actual := addTwoNumbers(test.input.l1, test.input.l2); !utils.SameList(actual, test.expect) {
			t.Errorf("input = %+v, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
