/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/11/18 4:11 下午
 */

package _sum_of_digits_of_string_after_convert

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		s string
		k int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{"iiii", 1}, 36},
		{input{"leetcode", 2}, 6},
		{input{"zbax", 2}, 8},
		{input{"hvmhoasabaymnmsd", 1}, 79},
	}

	for i, test := range tests {
		if actual := getLucky(test.input.s, test.input.k); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
