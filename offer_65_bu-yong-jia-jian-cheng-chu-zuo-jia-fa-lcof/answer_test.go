/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/29 3:11 下午
 */

package offer_65_bu_yong_jia_jian_cheng_chu_zuo_jia_fa_lcof

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		a, b int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{1, 2}, 3},
		{input{1, 1}, 2},
		{input{99, 1}, 100},
		{input{-99, 1}, -98},
	}

	for _, test := range tests {
		if actual := add(test.input.a, test.input.b); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
