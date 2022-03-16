/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/16 13:39
 */

package _number_of_1_bits

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  uint32
		expect int
	}{
		{11, 3},
		{128, 1},
		{4294967293, 31},
		{4294967295, 32},
	}

	for i, test := range tests {
		if actual := hammingWeight(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
