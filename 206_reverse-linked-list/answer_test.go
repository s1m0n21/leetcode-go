/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/18 12:09 下午
 */

package _reverse_linked_list

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct{
		input *utils.ListNode
		expect *utils.ListNode
	}{
		{utils.MakeListNode(1,2,3,4,5), utils.MakeListNode(5,4,3,2,1)},
		{utils.MakeListNode(9,8,7,6,5,4,3,2,1), utils.MakeListNode(1,2,3,4,5,6,7,8,9)},
		{nil, nil},
		{utils.MakeListNode(1), utils.MakeListNode(1)},
	}

	for _, test := range tests {
		if actual := reverseList(test.input); !utils.SameList(test.expect, actual) {
			t.Errorf("input = %s, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}

}
