/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 10:27 下午
 */

package _longest_substring_without_repeating_characters

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"hello", 3},
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, test := range tests {
		if actual := lengthOfLongestSubstring(test.input); actual != test.expect {
			t.Errorf("input = %s, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
