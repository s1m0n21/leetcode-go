/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/16 12:11 PM
 */

package _reverse_bits

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  uint32
		expect uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for i, test := range tests {
		if actual := reverseBits(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
