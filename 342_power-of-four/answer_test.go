/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/17 19:40
 */

package _power_of_four

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect bool
	}{
		{16, true},
		{5, false},
		{1, true},
	}

	for i, test := range tests {
		if actual := isPowerOfFour(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
