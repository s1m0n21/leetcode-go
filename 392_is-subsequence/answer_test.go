/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/31 2:10 上午
 */

package _is_subsequence

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		s, t string
	}

	tests := []struct {
		input  input
		expect bool
	}{
		{input{"abc", "ahbgdc"}, true},
		{input{"axc", "ahbgdc"}, false},
		{input{"ace", "abcde"}, true},
		{input{"aec", "abcde"}, false},
		{input{"", "abcde"}, true},
		{input{"b", "abc"}, true},
	}

	for _, test := range tests {
		if actual := isSubsequence(test.input.s, test.input.t); actual != test.expect {
			t.Errorf("input = %+v, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
