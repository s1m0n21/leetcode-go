/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/26 9:19 上午
 */

package _sum_of_two_integers

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		a, b int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{1, 1}, 1 + 1},
		{input{1, 2}, 1 + 2},
		{input{-1, 1}, -1 + 1},
	}

	for _, test := range tests {
		if actual := getSum(test.input.a, test.input.b); actual != test.expect {
			t.Errorf("input = %+v, expect = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
