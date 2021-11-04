/*
 * @Author       : s1m0n21
 * @Description  : Test answer
 * @Date         : 2021/11/04 12:12 PM
 */
package _valid_perfect_square

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect bool
	}{
		{16, true},
		{14, false},
		{1<<31 - 1, false},
	}

	for i, test := range tests {
		if actual := isPerfectSquare(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
