/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/2 9:06 上午
 */

package offer_22_lian_biao_zhong_dao_shu_di_kge_jie_dian_lcof

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		head *utils.ListNode
		k    int
	}

	list1 := utils.MakeListNode(1, 2, 3, 4, 5)
	list2 := utils.MakeListNode(1, 2, 3, 4, 5, 6)

	tests := []struct {
		input  input
		expect *utils.ListNode
	}{
		{input{list1, 2}, list1.Get(4)},
		{input{list2, 3}, list2.Get(4)},
	}

	for _, test := range tests {
		h := test.input.head
		if actual := getKthFromEnd(test.input.head, test.input.k); actual != test.expect {
			t.Errorf("input = (head = %s, k = %d), expect = %s, actual = %s",
				h, test.input.k, test.expect, actual,
				)
		}
	}
}
