/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 11:21 上午
 */

package _palindrome_number

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input int
		expect bool
	}{
		{-121, false},
		{121, true},
		{0, true},
		{10, false},
		{101, true},
	}

	for _, test := range tests {
		if actual := isPalindrome(test.input); actual != test.expect {
			t.Errorf("input = %d, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
