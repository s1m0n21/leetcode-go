/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/14 11:17 AM
 */

package _sum_of_digits_in_base_k

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		n, k int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{34, 6}, 9},
		{input{10, 10}, 1},
	}

	for i, test := range tests {
		if actual := sumBase(test.input.n, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
