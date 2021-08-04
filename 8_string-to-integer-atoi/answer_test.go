/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/5 2:21 上午
 */

package _string_to_integer_atoi

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+-12", 0},
		{"00000-42a1234", 0},
		{"+1", 1},
		{"  -0012a42", -12},
		{"9223372036854775808", 2147483647},
		{"-2147483648", -2147483648},
	}

	for _, test := range tests {
		if actual := myAtoi(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
