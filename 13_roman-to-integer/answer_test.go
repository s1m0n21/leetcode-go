/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 10:57 上午
 */

package _roman_to_integer

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input string
		expect int
	}{
		{"MCMXCIV", 1994},
		{"IV", 4},
		{"IXV", 14},
	}

	for _, test := range tests {
		if actual := romanToInt(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
