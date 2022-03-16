/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/17 01:35
 */

package _power_of_two

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect bool
	}{
		{1, true},
		{16, true},
		{3, false},
		{4, true},
		{5, false},
	}

	for i, test := range tests {
		if actual := isPowerOfTwo(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
