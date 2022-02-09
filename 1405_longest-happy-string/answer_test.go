/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/8 10:33 AM
 */

package _longest_happy_string

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		a, b, c int
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{1, 1, 7}, "ccaccbcc"},
		{input{2, 2, 1}, "abbac"},
		{input{7, 1, 0}, "aabaa"},
	}

	for i, test := range tests {
		if actual := longestDiverseString(test.input.a, test.input.b, test.input.c); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
