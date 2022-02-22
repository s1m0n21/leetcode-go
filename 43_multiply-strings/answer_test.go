/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/22 10:50 AM
 */

package _multiply_strings

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		num1, num2 string
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{"2", "3"}, "6"},
		{input{"0", "3"}, "0"},
		{input{"123", "456"}, "56088"},
	}

	for i, test := range tests {
		if actual := multiply(test.input.num1, test.input.num2); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
