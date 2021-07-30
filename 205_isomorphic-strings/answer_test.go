/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/4 12:17 上午
 */

package _isomorphic_strings

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		s string
		t string
	}

	tests := []struct{
		input input
		expect bool
	}{
		{input{"egg", "add"}, true},
		{input{"foo", "bar"}, false},
		{input{"paper", "title"}, true},
		{input{"123", "1234"}, false},
		{input{"", ""}, true},
	}

	for _, test := range tests {
		if actual := isIsomorphic(test.input.s, test.input.t); actual != test.expect {
			t.Errorf("input = %+v, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
