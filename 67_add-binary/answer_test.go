/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/24 5:50 下午
 */

package _add_binary

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		a, b string
	}

	tests := []struct {
		input  input
		expect string
	}{
		{input{"11", "1"}, "100"},
		{input{"1010", "1011"}, "10101"},
	}

	for _, test := range tests {
		if actual := addBinary(test.input.a, test.input.b); actual != test.expect {
			t.Errorf("input = %+v, expect = %s, actual = %s", test.input, test.expect, actual)
		}
	}
}
