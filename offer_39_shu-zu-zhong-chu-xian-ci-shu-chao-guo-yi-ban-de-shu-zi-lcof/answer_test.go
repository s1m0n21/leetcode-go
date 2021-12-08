/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/12/8 7:11 PM
 */

package offer_39_shu_zu_zhong_chu_xian_ci_shu_chao_guo_yi_ban_de_shu_zi_lcof

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 2, 2, 2, 5, 4, 2}, 2},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for i, test := range tests {
		if actual := majorityElement(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
