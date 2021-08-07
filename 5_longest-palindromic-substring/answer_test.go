/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/6 3:48 下午
 */

package _longest_palindromic_substring

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"aba", "aba"},
	}

	for _, test := range tests {
		if actual := longestPalindrome(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
