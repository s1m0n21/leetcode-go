/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/11 3:38 下午
 */

package _thousand_separator

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input int
		expect string
	}{
		{9876, "9.876"},
		{123456789, "123.456.789"},
		{0, "0"},
		{1, "1"},
	}

	for i, test := range tests {
		if actual := thousandSeparator(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
