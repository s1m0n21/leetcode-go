/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/15 10:58 AM
 */

package _add_strings

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		num1, num2 string
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{"11", "123"}, "134"},
		{input{"456", "77"}, "533"},
		{input{"0", "0"}, "0"},
		{input{"9", "1"}, "10"},
	}

	for i, test := range tests {
		if actual := addStrings(test.input.num1, test.input.num2); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
