/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/20 4:44 下午
 */

package _reverse_string_ii

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		s string
		k int
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{"abcd", 2}, "bacd"},
		{input{"abcdefg", 2}, "bacdfeg"},
		{input{"abc", 4}, "cba"},
	}

	for _, test := range tests {
		if actual := reverseStr(test.input.s, test.input.k); actual != test.expect {
			t.Errorf("input = %+v, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
