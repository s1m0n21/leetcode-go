/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/15 11:28 AM
 */

package _longest_palindrome

import "testing"

func TestANswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
		{"bbb", 3},
	}

	for i, test := range tests {
		if actual := longestPalindrome(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
