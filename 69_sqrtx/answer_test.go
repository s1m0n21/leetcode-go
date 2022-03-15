/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/15 9:41 PM
 */

package _sqrtx

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{9, 3},
		{4, 2},
		{81, 9},
		{8, 2},
		{1, 1},
	}

	for i, test := range tests {
		if actual := mySqrt(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
