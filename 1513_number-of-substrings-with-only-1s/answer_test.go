/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/12 5:01 下午
 */

package _number_of_substrings_with_only_1s

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"0110111", 9},
		{"101", 2},
		{"111111", 21},
		{"000", 0},
	}

	for i, test := range tests {
		if actual := numSub(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
