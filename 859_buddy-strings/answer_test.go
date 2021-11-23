/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/23 10:38 上午
 */

package _buddy_strings

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		s, goal string
	}

	tests := []struct {
		input  input
		expect bool
	}{
		{input{"ab", "bc"}, false},
		{input{"ab", "ba"}, true},
		{input{"abc", "bc"}, false},
		{input{"abc", "abc"}, false},
		{input{"abc", "cba"}, true},
		{input{"aa", "aa"}, true},
		{input{"aaaaaaabc", "aaaaaaacb"}, true},
		{input{"aaaaaaabc", "aaaaaaabd"}, false},
	}

	for i, test := range tests {
		if actual := buddyStrings(test.input.s, test.input.goal); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
