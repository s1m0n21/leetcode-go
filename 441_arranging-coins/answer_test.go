/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/10 10:28 下午
 */

package _arranging_coins

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input int
		expect int
	}{
		{5, 2},
		{8, 3},
		{6, 3},
	}

	for i, test := range tests {
		if actual := arrangeCoins(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
