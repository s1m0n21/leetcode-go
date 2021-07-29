/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 10:46 上午
 */

package _reverse_integer

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{123, 321},
		{-123, -321},
		{0, 0},
	}

	for _, test := range tests {
		if actual := reverse(test.input); actual != test.expect {
			t.Errorf("input = %d, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
