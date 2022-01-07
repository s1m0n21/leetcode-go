/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/8 4:06 AM
 */

package _longest_common_subsequence

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		text1, text2 string
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{"abcde", "ace"}, 3},
		{input{"abc", "abc"}, 3},
		{input{"abc", "def"}, 0},
	}

	for i, test := range tests {
		if actual := longestCommonSubsequence(test.input.text1, test.input.text2); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
