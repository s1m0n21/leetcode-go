/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/8 11:50 下午
 */

package _powx_n

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		x float64
		n int
	}

	tests := []struct {
		input  input
		expect float64
	}{
		{input{2, 10}, 1024},
		{input{2, -2}, 0.25},
		{input{2.1, 3}, 9.261000000000001},
		{input{0.00001, 2147483647}, 0},
	}

	for _, test := range tests {
		if actual := myPow(test.input.x, test.input.n); actual != test.expect {
			t.Errorf("input = %+v, expect = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
