/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 2:00 上午
 */

package _valid_palindrome

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct{
		input string
		expect bool
	}{
		{"race a car", false},
		{"www", true},
		{"t", true},
		{"taaaa", false},
		{"A man, a plan, a canal: Panama", true},
	}

	for _, test := range tests {
		if actual := isPalindrome(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
