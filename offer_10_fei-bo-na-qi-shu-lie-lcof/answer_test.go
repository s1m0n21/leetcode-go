/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/4 2:45 上午
 */

package offer_10_fei_bo_na_qi_shu_lie_lcof

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{2, 1},
		{5, 5},
		{7, 13},
		{95, 407059028},
	}

	for _, test := range tests {
		if actual := fib(test.input); actual != test.expect {
			t.Errorf("input = %d, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
